package chaincfg

import (
	"testing"
)

// FuzzGetChainParams fuzzes the GetChainParams function with various string inputs
func FuzzGetChainParams(f *testing.F) {
	addSeedCorpus(f)

	f.Fuzz(func(t *testing.T, network string) {
		params, err := GetChainParams(network)
		validateResult(t, network, params, err)
	})
}

// addSeedCorpus adds valid and edge case inputs to the fuzz test corpus
func addSeedCorpus(f *testing.F) {
	// Add seed corpus with valid network names
	validNetworks := []string{"mainnet", "testnet", "regtest", "stn", "teratestnet", "tstn"}
	for _, network := range validNetworks {
		f.Add(network)
	}

	// Add some edge cases
	edgeCases := []string{"", " ", "MAINNET", "invalid-network", "main net", "test-net", "🚀", "mainnet\x00", "mainnet\n"}
	for _, edge := range edgeCases {
		f.Add(edge)
	}
}

// validateResult checks that GetChainParams returns consistent results
func validateResult(t *testing.T, network string, params *Params, err error) {
	if err != nil {
		// If there's an error, params should be nil
		if params != nil {
			t.Errorf("GetChainParams(%q) returned non-nil params with error: %v", network, err)
		}

		return
	}

	// If there's no error, params should not be nil
	if params == nil {
		t.Errorf("GetChainParams(%q) returned nil params without error", network)

		return
	}

	// Verify that valid networks return expected params
	expectedParams := getExpectedParams(network)
	if expectedParams != nil && params != expectedParams {
		t.Errorf("GetChainParams(%q) returned unexpected params", network)
	}
}

// getExpectedParams returns the expected Params for known networks
func getExpectedParams(network string) *Params {
	switch network {
	case "mainnet":
		return &MainNetParams
	case "testnet":
		return &TestNetParams
	case "regtest":
		return &RegressionNetParams
	case "stn":
		return &StnParams
	case "teratestnet":
		return &TeraTestNetParams
	case "tstn":
		return &TeraScalingTestNetParams
	default:
		return nil
	}
}
