package main

import (
	"strings"
	"time"

	"github.com/bsv-blockchain/go-bt/v2/chainhash"
	bec "github.com/bsv-blockchain/go-sdk/primitives/ec"
)

// buildReport renders the solved genesis block as a human-readable report,
// including ready-to-paste go-chaincfg byte arrays.
func buildReport(privKey *bec.PrivateKey, pubUncompressed []byte, text string, value uint64,
	bits, timestamp, nonce uint32, coinbase []byte, merkleRoot, hash *chainhash.Hash,
	sigScript, pkScript []byte,
) string {
	var sb strings.Builder

	line := strings.Repeat("=", 78)

	writef(&sb, "\n%s\nGENESIS BLOCK SOLVED\n%s\n", line, line)

	writef(&sb, "\n--- Keys (SAVE THESE for records) ---\n")
	writef(&sb, "Private key (hex):   %x\n", privKey.Serialize())
	writef(&sb, "Private key (WIF):   %s\n", privKey.Wif())
	writef(&sb, "Public key (uncomp): %x\n", pubUncompressed)

	writef(&sb, "\n--- Coinbase ---\n")
	writef(&sb, "Coinbase text:       %q (%d bytes)\n", text, len([]byte(text)))
	writef(&sb, "Coinbase value:      %d satoshis\n", value)
	writef(&sb, "Coinbase txid:       %s\n", merkleRoot.String())
	writef(&sb, "Coinbase raw hex:    %x\n", coinbase)

	writef(&sb, "\n--- Header ---\n")
	writef(&sb, "Version:             1\n")
	writef(&sb, "HashPrevBlock:       %s\n", (&chainhash.Hash{}).String())
	writef(&sb, "HashMerkleRoot:      %s\n", merkleRoot.String())
	writef(&sb, "Timestamp:           %d  (%s)\n", timestamp, time.Unix(int64(timestamp), 0).UTC().Format(time.RFC3339))
	writef(&sb, "Bits:                %08x\n", bits)
	writef(&sb, "Nonce:               %d  (0x%08x)\n", nonce, nonce)
	writef(&sb, "Block hash:          %s\n", hash.String())

	fullBlock := append(buildHeader(merkleRoot, timestamp, bits, nonce), 0x01)
	fullBlock = append(fullBlock, coinbase...)

	writef(&sb, "\n--- Full raw block hex ---\n%x\n", fullBlock)

	writef(&sb, "\n--- go-chaincfg genesis.go snippet ---\n")
	writef(&sb, "var customGenesisCoinbaseSigScript = []byte{\n%s}\n", goByteArray(sigScript))
	writef(&sb, "var customGenesisCoinbasePkScript = []byte{\n%s}\n", goByteArray(pkScript))
	writef(&sb, "\nvar customGenesisHash = chainhash.Hash([chainhash.HashSize]byte{\n%s})\n", goByteArray(hash.CloneBytes()))
	writef(&sb, "\nvar customGenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{\n%s})\n", goByteArray(merkleRoot.CloneBytes()))
	writef(&sb, `
var customGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: customGenesisMerkleRoot,
		Timestamp:  time.Unix(%d, 0),
		Bits:       0x%08x,
		Nonce:      0x%08x,
	},
	Transactions: []*wire.MsgTx{&customGenesisCoinbaseTx},
}
`, timestamp, bits, nonce)

	writef(&sb, "\n%s\n", line)

	return sb.String()
}

// goByteArray formats a byte slice as indented Go source (8 bytes per line).
func goByteArray(b []byte) string {
	var sb strings.Builder

	for i, x := range b {
		if i%8 == 0 {
			sb.WriteString("\t")
		}

		writef(&sb, "0x%02x,", x)

		if i%8 == 7 || i == len(b)-1 {
			sb.WriteString("\n")
		} else {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}
