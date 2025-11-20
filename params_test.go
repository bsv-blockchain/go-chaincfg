// Copyright (c) 2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"testing"

	"github.com/bsv-blockchain/go-wire"
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

// TestGetChainParamsBase tests GetChainParams for all supported and unsupported networks.
func TestGetChainParamsBase(t *testing.T) {
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

// TestValidPubKeyHashAddrID tests the IsPubKeyHashAddrID function for valid address IDs.
func TestValidPubKeyHashAddrID(t *testing.T) {
	assert.True(t, IsPubKeyHashAddrID(wire.MainNet, MainNetParams.LegacyPubKeyHashAddrID), "Expected valid PubKeyHashAddrID for MainNet")
	assert.True(t, IsPubKeyHashAddrID(wire.TestNet, TestNetParams.LegacyPubKeyHashAddrID), "Expected valid PubKeyHashAddrID for TestNet")
	assert.True(t, IsPubKeyHashAddrID(wire.RegTestNet, RegressionNetParams.LegacyPubKeyHashAddrID), "Expected valid PubKeyHashAddrID for RegressionNet")
	assert.True(t, IsPubKeyHashAddrID(wire.STN, StnParams.LegacyPubKeyHashAddrID), "Expected valid PubKeyHashAddrID for StnNet")
	assert.True(t, IsPubKeyHashAddrID(wire.TeraTestNet, TeraTestNetParams.LegacyPubKeyHashAddrID), "Expected valid PubKeyHashAddrID for TeraTestNet")
	assert.True(t, IsPubKeyHashAddrID(wire.TeraScalingTestNet, TeraScalingTestNetParams.LegacyPubKeyHashAddrID), "Expected valid PubKeyHashAddrID for TeraScalingTestNet")
	assert.False(t, IsPubKeyHashAddrID(wire.BitcoinNet(0), MainNetParams.LegacyPubKeyHashAddrID), "Expected valid PubKeyHashAddrID for unknown network with MainNet ID")
	assert.False(t, IsPubKeyHashAddrID(wire.BitcoinNet(999), 0x7F), "Expected valid PubKeyHashAddrID for unsupported network with custom ID")
}

// TestInvalidPubKeyHashAddrID tests the IsPubKeyHashAddrID function for invalid address IDs across different networks.
func TestInvalidPubKeyHashAddrID(t *testing.T) {
	assert.False(t, IsPubKeyHashAddrID(wire.MainNet, 0xFF), "Expected invalid PubKeyHashAddrID for MainNet")
	assert.False(t, IsPubKeyHashAddrID(wire.TestNet, 0x00), "Expected invalid PubKeyHashAddrID for TestNet")
	assert.False(t, IsPubKeyHashAddrID(wire.BitcoinNet(0), MainNetParams.LegacyPubKeyHashAddrID), "Expected invalid PubKeyHashAddrID for unknown network")
	assert.False(t, IsPubKeyHashAddrID(wire.BitcoinNet(999), 0x7F), "Expected invalid PubKeyHashAddrID for unsupported network")
}

// TestGetChainParams tests the GetChainParams function for various scenarios.
func TestGetChainParams(t *testing.T) {
	tests := []struct {
		name        string
		network     string
		expectError bool
		expected    *Params
	}{
		{
			name:        "Known network - mainnet",
			network:     "mainnet",
			expectError: false,
			expected:    &MainNetParams,
		},
		{
			name:        "Unknown network",
			network:     "unknown",
			expectError: true,
			expected:    nil,
		},
		{
			name:        "Empty network string",
			network:     "",
			expectError: true,
			expected:    nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			params, err := GetChainParams(tc.network)

			if tc.expectError {
				require.Error(t, err)
				require.Nil(t, params)
			} else {
				require.NoError(t, err)
				require.NotNil(t, params)
				require.Equal(t, tc.expected, params)
			}
		})
	}
}
