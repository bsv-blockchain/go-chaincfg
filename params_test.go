// Copyright (c) 2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestInvalidHashStr ensures the newShaHashFromStr function panics when used to
// with an invalid hash string.
func TestInvalidHashStr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid hash, got nil")
		}
	}()
	newHashFromStr("banana")
}

// TestSeeds ensures the right seeds are defined.
func TestSeeds(t *testing.T) {

	expectedSeeds := []DNSSeed{
		{"seed.bitcoinsv.io", true},
	}

	require.NotNil(t, MainNetParams.DNSSeeds, "Seed values are not set")
	require.Len(t, expectedSeeds, len(MainNetParams.DNSSeeds), "Incorrect number of seed values")
	assert.Equal(t, expectedSeeds, MainNetParams.DNSSeeds, "Seed values are incorrect")
}

// TestInternalMap ensures that the internalParamMapByAddrID is correctly populated
func TestInternalMap(t *testing.T) {
	// mainnet - legacyPubKeyHashAddrID
	_, ok := internalParamMapByAddrID[MainNetParams.LegacyPubKeyHashAddrID]
	require.True(t, ok, "MainNetParams should be registered in internalParamMapByAddrID")
	require.Same(t, &MainNetParams, internalParamMapByAddrID[MainNetParams.LegacyPubKeyHashAddrID], "Expected MainNetParams for LegacyPubKeyHashAddrID")

	// mainnet - legacyScriptHashAddrID
	_, ok = internalParamMapByAddrID[MainNetParams.LegacyScriptHashAddrID]
	require.True(t, ok, "MainNetParams should be registered in internalParamMapByAddrID")
	require.Same(t, &MainNetParams, internalParamMapByAddrID[MainNetParams.LegacyScriptHashAddrID], "Expected MainNetParams for LegacyScriptHashAddrID")

	// testnet - legacyPubKeyHashAddrID
	_, ok = internalParamMapByAddrID[TestNetParams.LegacyPubKeyHashAddrID]
	require.True(t, ok, "TestNetParams should be registered in internalParamMapByAddrID")
	require.Same(t, &TestNetParams, internalParamMapByAddrID[TestNetParams.LegacyPubKeyHashAddrID], "Expected TestNetParams for LegacyPubKeyHashAddrID")

	// testnet - legacyScriptHashAddrID
	_, ok = internalParamMapByAddrID[TestNetParams.LegacyScriptHashAddrID]
	require.True(t, ok, "TestNetParams should be registered in internalParamMapByAddrID")
	require.Same(t, &TestNetParams, internalParamMapByAddrID[TestNetParams.LegacyScriptHashAddrID], "Expected TestNetParams for LegacyScriptHashAddrID")

	// regressionnet - legacyPubKeyHashAddrID
	_, ok = internalParamMapByAddrID[RegressionNetParams.LegacyPubKeyHashAddrID]
	require.True(t, ok, "RegressionNetParams should be registered in internalParamMapByAddrID")
	require.Same(t, &RegressionNetParams, internalParamMapByAddrID[RegressionNetParams.LegacyPubKeyHashAddrID], "Expected RegressionNetParams for LegacyPubKeyHashAddrID")

	// regressionnet - legacyScriptHashAddrID
	_, ok = internalParamMapByAddrID[RegressionNetParams.LegacyScriptHashAddrID]
	require.True(t, ok, "RegressionNetParams should be registered in internalParamMapByAddrID")
	require.Same(t, &RegressionNetParams, internalParamMapByAddrID[RegressionNetParams.LegacyScriptHashAddrID], "Expected RegressionNetParams for LegacyScriptHashAddrID")

	// stn - legacyPubKeyHashAddrID
	_, ok = internalParamMapByAddrID[StnParams.LegacyPubKeyHashAddrID]
	require.True(t, ok, "StnParams should be registered in internalParamMapByAddrID")
	require.Same(t, &StnParams, internalParamMapByAddrID[StnParams.LegacyPubKeyHashAddrID], "Expected StnParams for LegacyPubKeyHashAddrID")

	// stn - legacyScriptHashAddrID
	_, ok = internalParamMapByAddrID[StnParams.LegacyScriptHashAddrID]
	require.True(t, ok, "StnParams should be registered in internalParamMapByAddrID")
	require.Same(t, &StnParams, internalParamMapByAddrID[StnParams.LegacyScriptHashAddrID], "Expected StnParams for LegacyScriptHashAddrID")

	// teratestnet - legacyPubKeyHashAddrID
	_, ok = internalParamMapByAddrID[TeraTestNetParams.LegacyPubKeyHashAddrID]
	require.True(t, ok, "TeraTestNetParams should be registered in internalParamMapByAddrID")
	require.Same(t, &TeraTestNetParams, internalParamMapByAddrID[TeraTestNetParams.LegacyPubKeyHashAddrID], "Expected TeraTestNetParams for LegacyPubKeyHashAddrID")

	// teratestnet - legacyScriptHashAddrID
	_, ok = internalParamMapByAddrID[TeraTestNetParams.LegacyScriptHashAddrID]
	require.True(t, ok, "TeraTestNetParams should be registered in internalParamMapByAddrID")
	require.Same(t, &TeraTestNetParams, internalParamMapByAddrID[TeraTestNetParams.LegacyScriptHashAddrID], "Expected TeraTestNetParams for LegacyScriptHashAddrID")

	// tstn - legacyPubKeyHashAddrID
	_, ok = internalParamMapByAddrID[TeraScalingTestNetParams.LegacyPubKeyHashAddrID]
	require.True(t, ok, "TeraScalingTestNetParams should be registered in internalParamMapByAddrID")
	require.Same(t, &TeraScalingTestNetParams, internalParamMapByAddrID[TeraScalingTestNetParams.LegacyPubKeyHashAddrID], "Expected TeraScalingTestNetParams for LegacyPubKeyHashAddrID")

	// tstn - legacyScriptHashAddrID
	_, ok = internalParamMapByAddrID[TeraScalingTestNetParams.LegacyScriptHashAddrID]
	require.True(t, ok, "TeraScalingTestNetParams should be registered in internalParamMapByAddrID")
	require.Same(t, &TeraScalingTestNetParams, internalParamMapByAddrID[TeraScalingTestNetParams.LegacyScriptHashAddrID], "Expected TeraScalingTestNetParams for LegacyScriptHashAddrID")
}

// TestGetChainParams tests GetChainParams for all supported and unsupported networks.
func TestGetChainParams(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		wantPtr *Params
	}{
		{"mainnet", "mainnet", false, &MainNetParams},
		{"testnet", "testnet", false, &TestNetParams},
		{"regtest", "regtest", false, &RegressionNetParams},
		{"stn", "stn", false, &StnParams},
		{"teratestnet", "teratestnet", false, &TeraTestNetParams},
		{"tstn", "tstn", false, &TeraScalingTestNetParams},
		{"unknown", "unknown", true, nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetChainParams(tc.input)
			if tc.wantErr {
				require.Error(t, err, "expected error for input %q", tc.input)
				assert.Nil(t, got, "expected nil result for input %q", tc.input)
			} else {
				require.NoError(t, err, "unexpected error for input %q", tc.input)
				assert.Equal(t, tc.wantPtr, got, "expected pointer for input %q", tc.input)
			}
		})
	}
}

// TestDNSSeedString tests the String method of the DNSSeed type.
func TestDNSSeedString(t *testing.T) {
	seed := DNSSeed{Host: "example.com", HasFiltering: true}
	assert.Equal(t, "example.com", seed.String(), "DNSSeed.String() should return the Host field")
}
