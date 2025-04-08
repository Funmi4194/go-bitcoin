package bitcoin

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/stretchr/testify/assert"
)

// TestPubKeyFromString will test the method PubKeyFromString()
func TestPubKeyFromString(t *testing.T) {

	t.Parallel()

	var tests = []struct {
		inputKey       string
		expectedPubKey string
		expectedNil    bool
		expectedError  bool
	}{
		{"", "", true, true},
		{"0", "", true, true},
		{"00000", "", true, true},
		{"0309a1ede55bcb4d7ecbf45f015ea8e2f43cd71be97291d314f3be6871733f541b", "0309a1ede55bcb4d7ecbf45f015ea8e2f43cd71be97291d314f3be6871733f541b", false, false},
		{"022d35c7ede60cb68dee6e60ab9ad5863a7a726b297273d7a99b9dfb032a10e3f8", "022d35c7ede60cb68dee6e60ab9ad5863a7a726b297273d7a99b9dfb032a10e3f8", false, false},
	}

	for _, test := range tests {
		if pubKey, err := PubKeyFromString(test.inputKey); err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error not expected but got: %s", t.Name(), test.inputKey, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error was expected", t.Name(), test.inputKey)
		} else if pubKey != nil && test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and nil was expected", t.Name(), test.inputKey)
		} else if pubKey == nil && !test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and nil was NOT expected", t.Name(), test.inputKey)
		} else if pubKey != nil && hex.EncodeToString(pubKey.SerializeCompressed()) != test.expectedPubKey {
			t.Fatalf("%s Failed: [%s] inputted and [%s] expected, but got: %s", t.Name(), test.inputKey, test.expectedPubKey, hex.EncodeToString(pubKey.SerializeCompressed()))
		}
	}
}

// TestPubKeyFromPrivateKey will test the method PubKeyFromPrivateKey()
func TestPubKeyFromPrivateKey(t *testing.T) {
	t.Parallel()

	priv, err := PrivateKeyFromString("fff9f5137145b7609070fcaf13ab2db3974230699c74b1a3ca5479fb506b5de9")
	assert.NoError(t, err)
	assert.NotNil(t, priv)

	var tests = []struct {
		inputKey       *btcec.PrivateKey
		expectedPubKey string
	}{
		{priv, "022d35c7ede60cb68dee6e60ab9ad5863a7a726b297273d7a99b9dfb032a10e3f8"},
	}

	for _, test := range tests {
		if pubKey := PubKeyFromPrivateKey(test.inputKey); pubKey != test.expectedPubKey {
			t.Fatalf("%s Failed: [%v] inputted and [%s] expected, but got: %s", t.Name(), test.inputKey, test.expectedPubKey, pubKey)
		}
	}
}

// TestPubKeyFromPrivateKeyString will test the method PubKeyFromPrivateKeyString()
func TestPubKeyFromPrivateKeyString(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		inputKey       string
		expectedPubKey string
		expectedError  bool
	}{
		{"fff9f5137145b7609070fcaf13ab2db3974230699c74b1a3ca5479fb506b5de9", "022d35c7ede60cb68dee6e60ab9ad5863a7a726b297273d7a99b9dfb032a10e3f8", false},
		{"54035dd4c7dda99ac473905a3d82f7864322b49bab1ff441cc457183b9bd8ab", "", true},
		{"0", "", true},
		{"", "", true},
	}

	for _, test := range tests {
		if pubKey, err := PubKeyFromPrivateKeyString(test.inputKey); err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error not expected but got: %s", t.Name(), test.inputKey, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error was expected", t.Name(), test.inputKey)
		} else if pubKey != test.expectedPubKey {
			t.Fatalf("%s Failed: [%s] inputted and [%s] expected, but got: %s", t.Name(), test.inputKey, test.expectedPubKey, pubKey)
		}
	}
}
