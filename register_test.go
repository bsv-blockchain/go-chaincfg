package chaincfg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define some of the required parameters for a user-registered
// network.  This is necessary to test the registration of and
// lookup of encoding magics from the network.
var mockNetParams = Params{
	Name: "mocknet",
	Net:  1<<32 - 1,

	LegacyPubKeyHashAddrID: 0x9f,
	LegacyScriptHashAddrID: 0xf9,
	HDPrivateKeyID:         [4]byte{0x01, 0x02, 0x03, 0x04},
	HDPublicKeyID:          [4]byte{0x05, 0x06, 0x07, 0x08},
	CashAddressPrefix:      "bsvmock",
}

// TestSuite is a struct that embeds testify's suite.Suite to create a test suite for the chaincfg package.
type TestSuite struct {
	suite.Suite
}

// TestRegisterFlow tests the registration flow of networks, ensuring that
func (ts *TestSuite) TestRegisterFlow() {
	builtins := []*Params{&MainNetParams, &RegressionNetParams, &TestNetParams}

	// 1. Built-in nets should already resolve magics/HD before explicit Register().
	ts.Run("baseline-builtins", func() {
		for _, p := range builtins {
			ts.assertAddrMagics(p, true)
			ts.assertHD(p, false)
		}
	})

	// 2. Register built-ins (should succeed) and then ensure duplicates fail.
	ts.Run("register-builtins", func() {
		for _, p := range builtins {
			ts.Require().NoError(Register(p), "first register %s", p.Name)
			ts.Require().ErrorIs(Register(p), ErrDuplicateNet, "duplicate register %s", p.Name)
		}
	})

	// 3. mocknet flow: invalid → register → valid.
	ts.Run("mocknet-flow", func() {
		ts.assertAddrMagics(&mockNetParams, false)
		ts.assertHD(&mockNetParams, true)

		ts.Require().NoError(Register(&mockNetParams))
		ts.Require().ErrorIs(Register(&mockNetParams), ErrDuplicateNet)

		ts.assertAddrMagics(&mockNetParams, true)
		ts.assertHD(&mockNetParams, false)
	})

	// 4. Edge-case invalid inputs preserved from original tests.
	ts.Run("invalid-edge-cases", func() {
		ts.Require().False(IsPubKeyHashAddrID(MainNetParams.Net, 0xff))
		ts.Require().False(IsScriptHashAddrID(MainNetParams.Net, 0xff))
		ts.Require().False(IsCashAddressPrefix(MainNetParams.Net, "abc1"))
		ts.Require().False(IsCashAddressPrefix(MainNetParams.Net, "1"))
		ts.Require().False(IsCashAddressPrefix(MainNetParams.Net, MainNetParams.CashAddressPrefix))

		_, err := HDPrivateKeyToPublicKeyID([]byte{0xff, 0xff, 0xff, 0xff})
		ts.Require().ErrorIs(err, ErrUnknownHDKeyID)

		_, err = HDPrivateKeyToPublicKeyID([]byte{0xff})
		ts.Require().ErrorIs(err, ErrUnknownHDKeyID)
	})

	// 5. Final duplicate sweep for *all* registered nets.
	ts.Run("duplicate-all-nets", func() {
		all := append(builtins, &mockNetParams)
		for _, p := range all {
			ts.Require().ErrorIs(Register(p), ErrDuplicateNet, "duplicate final %s", p.Name)
		}
	})
}

// assertAddrMagics checks if the address magic IDs and cash address prefixes
func (ts *TestSuite) assertAddrMagics(p *Params, want bool) {
	ts.Equal(want, IsPubKeyHashAddrID(p.Net, p.LegacyPubKeyHashAddrID), "P2PKH magic %s", p.Name)
	ts.Equal(want, IsScriptHashAddrID(p.Net, p.LegacyScriptHashAddrID), "P2SH magic %s", p.Name)

	full := p.CashAddressPrefix + ":"
	ts.Equal(want, IsCashAddressPrefix(p.Net, full), "cashaddr %s", p.Name)
	ts.Equal(want, IsCashAddressPrefix(p.Net, strings.ToUpper(full)), "cashaddr upper %s", p.Name)
}

// assertHD checks if the HDPrivateKeyID can be converted to a HDPublicKeyID.
func (ts *TestSuite) assertHD(p *Params, wantErr bool) {
	pub, err := HDPrivateKeyToPublicKeyID(p.HDPrivateKeyID[:])
	if wantErr {
		ts.ErrorIs(err, ErrUnknownHDKeyID, "HD priv->pub should fail for %s", p.Name)
	} else {
		ts.Require().NoError(err, "HD priv->pub failed for %s", p.Name)
		ts.Equal(p.HDPublicKeyID[:], pub, "HD pub mismatch for %s", p.Name)
	}
}

// TestRegisterSuite runs the test suite for the chaincfg package.
func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
