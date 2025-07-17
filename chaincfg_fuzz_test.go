package chaincfg

import (
	"testing"
)

// FuzzGetChainParams fuzzes the GetChainParams function with various string inputs
func FuzzGetChainParams(f *testing.F) {
	// Add seed corpus with valid network names
	f.Add("mainnet")
	f.Add("testnet")
	f.Add("regtest")
	f.Add("stn")
	f.Add("teratestnet")
	f.Add("tstn")
	
	// Add some edge cases
	f.Add("")
	f.Add(" ")
	f.Add("MAINNET")
	f.Add("invalid-network")
	f.Add("main net")
	f.Add("test-net")
	f.Add("🚀")
	f.Add("mainnet\x00")
	f.Add("mainnet\n")
	
	f.Fuzz(func(t *testing.T, network string) {
		params, err := GetChainParams(network)
		
		// Check that we either get a valid result or an error
		if err != nil {
			// If there's an error, params should be nil
			if params != nil {
				t.Errorf("GetChainParams(%q) returned non-nil params with error: %v", network, err)
			}
		} else {
			// If there's no error, params should not be nil
			if params == nil {
				t.Errorf("GetChainParams(%q) returned nil params without error", network)
			}
			
			// Verify that valid networks return expected params
			switch network {
			case "mainnet":
				if params != &MainNetParams {
					t.Errorf("GetChainParams(%q) returned unexpected params", network)
				}
			case "testnet":
				if params != &TestNetParams {
					t.Errorf("GetChainParams(%q) returned unexpected params", network)
				}
			case "regtest":
				if params != &RegressionNetParams {
					t.Errorf("GetChainParams(%q) returned unexpected params", network)
				}
			case "stn":
				if params != &StnParams {
					t.Errorf("GetChainParams(%q) returned unexpected params", network)
				}
			case "teratestnet":
				if params != &TeraTestNetParams {
					t.Errorf("GetChainParams(%q) returned unexpected params", network)
				}
			case "tstn":
				if params != &TeraScalingTestNetParams {
					t.Errorf("GetChainParams(%q) returned unexpected params", network)
				}
			}
		}
	})
}
