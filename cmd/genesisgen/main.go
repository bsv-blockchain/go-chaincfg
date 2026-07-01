// Command genesisgen builds and mines a Bitcoin-style genesis block for a
// custom (e.g. test scaling) network.
//
// It mirrors the structure of Satoshi's original genesis block:
//   - the arbitrary headline is embedded in the coinbase input scriptSig,
//     preceded by a push of the difficulty bits (like the original 04ffff001d),
//   - the single coinbase output is a P2PK script (<pubkey> OP_CHECKSIG).
//
// The difficulty (nBits), coinbase text and timestamp are all settable. A fresh
// keypair is generated for provenance and its private key is printed.
package main

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bsv-blockchain/go-bt/v2"
	"github.com/bsv-blockchain/go-bt/v2/bscript"
	"github.com/bsv-blockchain/go-bt/v2/chainhash"
	bec "github.com/bsv-blockchain/go-sdk/primitives/ec"
)

const (
	difficultyOneBits = 0x1d00ffff
	headerSize        = 80
	headerPrefixSize  = 76
	nonceOffset       = 76
	minCoinbaseScript = 2
	maxCoinbaseScript = 100
	pubKeyPushLen     = 65
	nonceSpace        = uint64(1) << 32
	hashSize          = 32
)

var (
	errScriptSigSize = errors.New("coinbase scriptSig out of range (must be 2..100 bytes)")
	errZeroTarget    = errors.New("nBits encodes a zero target (unmineable)")
)

func main() {
	if err := run(); err != nil {
		writef(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	bitsStr := flag.String("bits", "1d00ffff", "difficulty target in compact nBits form (big-endian hex, e.g. 1d00ffff or 207fffff)")
	text := flag.String("text", "The Times 01/Jul/2026 Starmer puts Burnham in GBP5bn defence black hole", "arbitrary headline embedded in the coinbase scriptSig")
	timestamp := flag.Int64("timestamp", time.Now().Unix(), "block timestamp (unix seconds); may be bumped upward if no nonce solves the target")
	value := flag.Uint64("value", 5_000_000_000, "coinbase output value in satoshis (default 50 coins)")
	workers := flag.Int("workers", runtime.NumCPU(), "number of parallel mining goroutines")

	flag.Parse()

	bits64, err := strconv.ParseUint(*bitsStr, 16, 32)
	if err != nil {
		return fmt.Errorf("invalid -bits %q: %w", *bitsStr, err)
	}

	bits := uint32(bits64)

	target := compactToTarget(bits)
	if target.Sign() == 0 {
		return fmt.Errorf("invalid -bits %q: %w", *bitsStr, errZeroTarget)
	}

	// Fresh keypair for provenance. The genesis coinbase output is never inserted
	// into the UTXO set, so the key is printed rather than discarded.
	privKey, err := bec.NewPrivateKey()
	if err != nil {
		return fmt.Errorf("failed to generate private key: %w", err)
	}

	pubUncompressed := privKey.PubKey().Uncompressed()

	coinbase, sigScript, pkScript, err := buildGenesisCoinbase(bitsLE(bits), *text, *value, pubUncompressed)
	if err != nil {
		return err
	}

	tx, err := bt.NewTxFromBytes(coinbase)
	if err != nil {
		return fmt.Errorf("coinbase failed to parse: %w", err)
	}

	// For a single-transaction block the merkle root is the coinbase txid.
	merkleRoot := tx.TxIDChainHash()
	targetBE := target.FillBytes(make([]byte, hashSize))

	ts := uint32(*timestamp) //nolint:gosec // unix seconds fit uint32 until year 2106

	writef(os.Stderr, "Mining genesis block with %d workers (bits=%08x, difficulty=%.4f)...\n",
		*workers, bits, difficulty(target))

	var (
		nonce uint32
		ok    bool
	)

	for {
		prefix := buildHeaderPrefix(merkleRoot, ts, bits)

		nonce, ok = mine(prefix, targetBE, *workers)
		if ok {
			break
		}

		ts++

		writef(os.Stderr, "  no solution in the 2^32 nonce space; bumping timestamp to %d\n", ts)
	}

	hash := blockHash(buildHeader(merkleRoot, ts, bits, nonce))

	report := buildReport(privKey, pubUncompressed, *text, *value, bits, ts, nonce, coinbase, merkleRoot, hash, sigScript, pkScript)

	if _, err := io.WriteString(os.Stdout, report); err != nil {
		return fmt.Errorf("failed to write report: %w", err)
	}

	return nil
}

// buildGenesisCoinbase assembles the raw coinbase transaction bytes, mirroring
// the layout of Satoshi's original genesis coinbase. It also returns the
// scriptSig and locking script for reporting.
func buildGenesisCoinbase(bits []byte, text string, value uint64, pubUncompressed []byte) (coinbase, sigScript, pkScript []byte, err error) {
	// scriptSig: push(bits) push(0x04) push(text) — same shape as 04ffff001d0104<text>.
	sigScript = make([]byte, 0, maxCoinbaseScript)
	sigScript = append(sigScript, pushData(bits)...)
	sigScript = append(sigScript, pushData([]byte{0x04})...)
	sigScript = append(sigScript, pushData([]byte(text))...)

	if len(sigScript) < minCoinbaseScript || len(sigScript) > maxCoinbaseScript {
		return nil, nil, nil, fmt.Errorf("got %d bytes (shorten -text): %w", len(sigScript), errScriptSigSize)
	}

	// P2PK locking script: <pubkey> OP_CHECKSIG.
	pkScript = append(pushData(pubUncompressed), bscript.OpCHECKSIG)

	cb := make([]byte, 0, 256)

	version := make([]byte, 4)
	binary.LittleEndian.PutUint32(version, 1)
	cb = append(cb, version...)

	cb = append(cb, 0x01)                   // input count
	cb = append(cb, make([]byte, 32)...)    // previous txid: all zeros
	cb = append(cb, 0xff, 0xff, 0xff, 0xff) // previous index: 0xffffffff
	cb = append(cb, bt.VarInt(uint64(len(sigScript))).Bytes()...)
	cb = append(cb, sigScript...)
	cb = append(cb, 0xff, 0xff, 0xff, 0xff) // sequence

	cb = append(cb, 0x01) // output count

	valBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valBytes, value)
	cb = append(cb, valBytes...)
	cb = append(cb, bt.VarInt(uint64(len(pkScript))).Bytes()...)
	cb = append(cb, pkScript...)

	cb = append(cb, make([]byte, 4)...) // locktime: 0

	return cb, sigScript, pkScript, nil
}

// pushData returns the given data prefixed with the minimal push opcode(s).
func pushData(data []byte) []byte {
	n := len(data)

	switch {
	case n < 0x4c:
		return append([]byte{byte(n)}, data...)
	case n <= 0xff:
		return append([]byte{0x4c, byte(n)}, data...)
	default:
		return append([]byte{0x4d, byte(n), byte(n >> 8)}, data...) //nolint:gosec // n bounded by caller (<=100 bytes), fits two length bytes
	}
}

// bitsLE returns the compact nBits as 4 little-endian bytes (header/scriptSig order).
func bitsLE(bits uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)

	return b
}

// compactToTarget converts compact nBits to the full 256-bit target, mirroring
// SVNode's arith_uint256::SetCompact. It returns zero for malformed encodings.
func compactToTarget(bits uint32) *big.Int {
	exponent := bits >> 24
	mantissa := bits & 0x007fffff

	if mantissa != 0 && bits&0x00800000 != 0 {
		return big.NewInt(0) // negative
	}

	target := big.NewInt(int64(mantissa))

	if exponent <= 3 {
		target.Rsh(target, uint(8*(3-exponent)))
	} else {
		target.Lsh(target, uint(8*(exponent-3)))
	}

	return target
}

// difficulty reports the ratio of the difficulty-1 target to the given target.
func difficulty(target *big.Int) float64 {
	one := compactToTarget(difficultyOneBits)

	ratio := new(big.Float).Quo(new(big.Float).SetInt(one), new(big.Float).SetInt(target))
	f, _ := ratio.Float64()

	return f
}

// buildHeaderPrefix returns the first 76 bytes of the block header (everything
// except the 4-byte nonce).
func buildHeaderPrefix(merkleRoot *chainhash.Hash, timestamp, bits uint32) []byte {
	prefix := make([]byte, 0, headerPrefixSize)

	version := make([]byte, 4)
	binary.LittleEndian.PutUint32(version, 1)
	prefix = append(prefix, version...)

	prefix = append(prefix, make([]byte, hashSize)...) // HashPrevBlock: all zeros
	prefix = append(prefix, merkleRoot.CloneBytes()...)

	ts := make([]byte, 4)
	binary.LittleEndian.PutUint32(ts, timestamp)
	prefix = append(prefix, ts...)

	prefix = append(prefix, bitsLE(bits)...)

	return prefix
}

// buildHeader returns the full 80-byte serialized block header.
func buildHeader(merkleRoot *chainhash.Hash, timestamp, bits, nonce uint32) []byte {
	header := make([]byte, headerSize)
	copy(header, buildHeaderPrefix(merkleRoot, timestamp, bits))
	binary.LittleEndian.PutUint32(header[nonceOffset:], nonce)

	return header
}

// blockHash returns the double-SHA256 of the header as a chainhash (internal order).
func blockHash(header []byte) *chainhash.Hash {
	first := sha256.Sum256(header)
	second := sha256.Sum256(first[:])
	h, _ := chainhash.NewHash(second[:])

	return h
}

// mine scans the entire 2^32 nonce space in parallel, returning the first nonce
// whose double-SHA256 header hash is <= target. ok is false if none exists.
func mine(prefix, targetBE []byte, workers int) (nonce uint32, ok bool) {
	if workers < 1 {
		workers = 1
	}

	var (
		found       int32
		resultNonce uint32
		wg          sync.WaitGroup
	)

	chunk := nonceSpace / uint64(workers)

	for w := range workers {
		start := uint64(w) * chunk

		end := start + chunk
		if w == workers-1 {
			end = nonceSpace
		}

		wg.Add(1)

		go func(start, end uint64) {
			defer wg.Done()

			if n, hit := scanRange(prefix, targetBE, start, end, &found); hit && atomic.CompareAndSwapInt32(&found, 0, 1) {
				resultNonce = n
			}
		}(start, end)
	}

	wg.Wait()

	return resultNonce, found != 0
}

// scanRange tries nonces in [start, end), returning the first that meets target.
func scanRange(prefix, targetBE []byte, start, end uint64, found *int32) (uint32, bool) {
	hdr := make([]byte, headerSize)
	copy(hdr, prefix)

	for n := start; n < end; n++ {
		if n&0xffff == 0 && atomic.LoadInt32(found) != 0 {
			return 0, false
		}

		binary.LittleEndian.PutUint32(hdr[nonceOffset:], uint32(n)) //nolint:gosec // n < 2^32 by loop bound

		first := sha256.Sum256(hdr)
		second := sha256.Sum256(first[:])

		if meetsTarget(&second, targetBE) {
			return uint32(n), true //nolint:gosec // n < 2^32 by loop bound
		}
	}

	return 0, false
}

// meetsTarget reports whether the block hash (the double-SHA256 result read as a
// little-endian 256-bit integer) is <= the big-endian target.
func meetsTarget(hash *[hashSize]byte, targetBE []byte) bool {
	for i := range hashSize {
		hb := hash[hashSize-1-i] // reverse: little-endian hash -> big-endian value
		tb := targetBE[i]

		if hb < tb {
			return true
		}

		if hb > tb {
			return false
		}
	}

	return true
}

// writef writes a formatted line to w, ignoring the (always nil for these
// destinations) write error.
func writef(w io.Writer, format string, a ...any) {
	_, _ = fmt.Fprintf(w, format, a...)
}
