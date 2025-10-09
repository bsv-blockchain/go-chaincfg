// Package examples provides example code for various functionalities.
package examples

import (
	"fmt"
	"os"

	chaincfg "github.com/bsv-blockchain/go-chaincfg"
	"github.com/bsv-blockchain/go-wire"
)

// ExampleIsPubKeyHashAddrID demonstrates how to verify the legacy public key hash
// (P2PKH) address identifier bytes for a specific Bitcoin SV network.
//
// The helper reports whether the provided identifier is valid for the supplied
// network magic value. This is useful when parsing external configuration or
// performing sanity checks on user-supplied address metadata.
func ExampleIsPubKeyHashAddrID() {
	write := func(value bool) {
		if _, err := fmt.Fprintf(os.Stdout, "%t\n", value); err != nil {
			panic(err)
		}
	}

	write(chaincfg.IsPubKeyHashAddrID(wire.MainNet, chaincfg.MainNetParams.LegacyPubKeyHashAddrID))
	write(chaincfg.IsPubKeyHashAddrID(wire.MainNet, chaincfg.TestNetParams.LegacyPubKeyHashAddrID))

	// Output:
	// true
	// false
}
