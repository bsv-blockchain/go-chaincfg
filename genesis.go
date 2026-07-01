// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/bsv-blockchain/go-bt/v2/chainhash"
	"github.com/bsv-blockchain/go-wire"
)

// GenesisActivationHeight is the block height at which the genesis block
var GenesisActivationHeight = uint32(620538)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
				0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
				0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
				0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
				0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
				0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
				0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec|*/
				0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
				0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for |*/
				0x62, 0x61, 0x6e, 0x6b, 0x73, /* |banks| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, /* |A.g....U| */
				0x48, 0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, /* |H'.g..q0| */
				0xb7, 0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, /* |..\..(.9| */
				0x09, 0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, /* |..yb...a| */
				0xde, 0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, /* |..I..?L.| */
				0x38, 0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, /* |8..U....| */
				0x12, 0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, /* |..\8M...| */
				0x8d, 0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, /* |.W.Lp+k.| */
				0x1d, 0x5f, 0xac, /* |._.| */
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the blockchain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x6f, 0xe2, 0x8c, 0x0a, 0xb6, 0xf1, 0xb3, 0x72,
	0xc1, 0xa6, 0xa2, 0x46, 0xae, 0x63, 0xf7, 0x4f,
	0x93, 0x1e, 0x83, 0x65, 0xe1, 0x5a, 0x08, 0x9c,
	0x68, 0xd6, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x3b, 0xa3, 0xed, 0xfd, 0x7a, 0x7b, 0x12, 0xb2,
	0x7a, 0xc7, 0x2c, 0x3e, 0x67, 0x76, 0x8f, 0x61,
	0x7f, 0xc8, 0x1b, 0xc3, 0x88, 0x8a, 0x51, 0x32,
	0x3a, 0x9f, 0xb8, 0xaa, 0x4b, 0x1e, 0x5e, 0x4a,
})

// genesisBlock defines the genesis block of the blockchain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(0x495fab29, 0), // 2009-01-03 18:15:05 +0000 UTC
		Bits:       0x1d00ffff,               // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x7c2bac1d,               // 2083236893
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the blockchain which serves
// as the public transaction ledger for the local regression test network.
// Please note the coinbase is the same as the mainnet genesis block, but the timestamp, nonce and
// difficulty are different.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(0x4d49e5da, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      0x2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNetGenesisHash is the hash of the first block in the blockchain for the
// test network (version 3).
var testNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x43, 0x49, 0x7f, 0xd7, 0xf8, 0x26, 0x95, 0x71,
	0x08, 0xf4, 0xa3, 0x0f, 0xd9, 0xce, 0xc3, 0xae,
	0xba, 0x79, 0x97, 0x20, 0x84, 0xe9, 0x0e, 0xad,
	0x01, 0xea, 0x33, 0x09, 0x00, 0x00, 0x00, 0x00,
})

// testNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNetGenesisMerkleRoot = genesisMerkleRoot

// testNetGenesisBlock defines the genesis block of the blockchain which
// serves as the public transaction ledger for the test network (version 3).
var testNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNetGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1d00ffff,               // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x18aea41a,               // 414098458
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// teraTestNetGenesisBlock defines the genesis block of the blockchain which
// serves as the public transaction ledger for the test network (version 3).
var teraTestNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNetGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1755606836, 0), // 2025-08-19T08:33:56 +0000 UTC
		Bits:       0x1d00ffff,
		Nonce:      0x411f6c9c, // Nonce value that produces a block hash meeting the proof-of-work requirement for 0x1d00ffff difficulty
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// teraTestNetGenesisHash is the hash of the first block in the blockchain for the
// tera test network.
var teraTestNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x6d, 0x77, 0xb7, 0x76, 0x79, 0x81, 0xea, 0xc2,
	0xb2, 0x04, 0x4a, 0x1a, 0x1c, 0x19, 0xb9, 0x74,
	0x1c, 0x23, 0x47, 0x37, 0x5b, 0x8f, 0xa8, 0xa0,
	0xbb, 0xea, 0x99, 0x04, 0x00, 0x00, 0x00, 0x00,
})

// stnGenesisHash is the hash of the first block in the blockchain for the
// stn network.
var stnGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x43, 0x49, 0x7f, 0xd7, 0xf8, 0x26, 0x95, 0x71,
	0x08, 0xf4, 0xa3, 0x0f, 0xd9, 0xce, 0xc3, 0xae,
	0xba, 0x79, 0x97, 0x20, 0x84, 0xe9, 0x0e, 0xad,
	0x01, 0xea, 0x33, 0x09, 0x00, 0x00, 0x00, 0x00,
})

// stnGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the stn network.  It is the same as the merkle root for the main
// network.
var stnGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the blockchain which serves
// as the public transaction ledger for the regression test network.
var stnGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: stnGenesisMerkleRoot,     // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1d00ffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      0x0a5bac18,               // 414098458
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// scalingTeraTestNetGenesisCoinbaseTx is the coinbase transaction for the
// scaling tera test network genesis block. Unlike the other networks it uses a
// unique coinbase: the headline "The Times 01/Jul/2026 Starmer puts Burnham in
// £5bn defence black hole" embedded in the input scriptSig, and a P2PK output
// paying a freshly generated (throwaway) key. The output is unspendable in
// practice because the genesis coinbase is never added to the UTXO set.
var scalingTeraTestNetGenesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x46, /* |.......F| */
				0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
				0x73, 0x20, 0x30, 0x31, 0x2f, 0x4a, 0x75, 0x6c, /* |s 01/Jul| */
				0x2f, 0x32, 0x30, 0x32, 0x36, 0x20, 0x53, 0x74, /* |/2026 St| */
				0x61, 0x72, 0x6d, 0x65, 0x72, 0x20, 0x70, 0x75, /* |armer pu| */
				0x74, 0x73, 0x20, 0x42, 0x75, 0x72, 0x6e, 0x68, /* |ts Burnh| */
				0x61, 0x6d, 0x20, 0x69, 0x6e, 0x20, 0xc2, 0xa3, /* |am in ..| */
				0x35, 0x62, 0x6e, 0x20, 0x64, 0x65, 0x66, 0x65, /* |5bn defe| */
				0x6e, 0x63, 0x65, 0x20, 0x62, 0x6c, 0x61, 0x63, /* |nce blac| */
				0x6b, 0x20, 0x68, 0x6f, 0x6c, 0x65, /* |k hole| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0xa2, 0x16, 0xe8, 0x98, 0x88, 0x7c,
				0x93, 0x5d, 0x15, 0xf9, 0xbe, 0x18, 0xfb, 0x99,
				0x41, 0x8b, 0xb7, 0xed, 0x0e, 0x19, 0x06, 0x92,
				0x72, 0x88, 0x14, 0xa1, 0xff, 0x75, 0x5c, 0xdf,
				0x59, 0x90, 0xce, 0x8a, 0x1a, 0xea, 0x3b, 0xdd,
				0x72, 0x0d, 0xe2, 0x18, 0xb9, 0x51, 0x55, 0xf2,
				0xb4, 0xae, 0xd5, 0xc6, 0x5e, 0x63, 0xbf, 0x72,
				0x64, 0xb7, 0x61, 0xb8, 0x7a, 0x27, 0xdc, 0x0d,
				0x6a, 0x73, 0xac,
			},
		},
	},
	LockTime: 0,
}

// scalingTeraTestNetGenesisMerkleRoot is the hash of the coinbase transaction in
// the scaling tera test network genesis block (txid
// 64452e5b25c65e492ad6a4f5ce9f427ca986626c28315d88de920d66e28cc98f).
var scalingTeraTestNetGenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x8f, 0xc9, 0x8c, 0xe2, 0x66, 0x0d, 0x92, 0xde,
	0x88, 0x5d, 0x31, 0x28, 0x6c, 0x62, 0x86, 0xa9,
	0x7c, 0x42, 0x9f, 0xce, 0xf5, 0xa4, 0xd6, 0x2a,
	0x49, 0x5e, 0xc6, 0x25, 0x5b, 0x2e, 0x45, 0x64,
})

// scalingTeraTestNetGenesisBlock defines the genesis block of the blockchain
// which serves as the public transaction ledger for the scaling tera test
// network.
var scalingTeraTestNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},                    // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: scalingTeraTestNetGenesisMerkleRoot, // 64452e5b25c65e492ad6a4f5ce9f427ca986626c28315d88de920d66e28cc98f
		Timestamp:  time.Unix(1782864000, 0),            // 2026-07-01 00:00:00 +0000 UTC
		Bits:       0x1d00ffff,                          // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x6a201818,                          // 1780488216
	},
	Transactions: []*wire.MsgTx{&scalingTeraTestNetGenesisCoinbaseTx},
}

// scalingTeraTestNetGenesisHash is the hash of the first block in the blockchain
// for the scaling tera test network
// (000000005d221c0e023cb56b5682cf094f32cd959958b40bc931e5797cae706c).
var scalingTeraTestNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x6c, 0x70, 0xae, 0x7c, 0x79, 0xe5, 0x31, 0xc9,
	0x0b, 0xb4, 0x58, 0x99, 0x95, 0xcd, 0x32, 0x4f,
	0x09, 0xcf, 0x82, 0x56, 0x6b, 0xb5, 0x3c, 0x02,
	0x0e, 0x1c, 0x22, 0x5d, 0x00, 0x00, 0x00, 0x00,
})
