package bitcoin

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPrivateKeyFromString will test the method PrivateKeyFromString()
func TestPrivateKeyFromString(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input         string
		expectedKey   string
		expectedNil   bool
		expectedError bool
	}{
		{"54035dd4c7dda99ac473905a3d82f7864322b49bab1ff441cc457183b9bd8abd", "54035dd4c7dda99ac473905a3d82f7864322b49bab1ff441cc457183b9bd8abd", false, false},
		{"fff9f5137145b7609070fcaf13ab2db3974230699c74b1a3ca5479fb506b5de9", "fff9f5137145b7609070fcaf13ab2db3974230699c74b1a3ca5479fb506b5de9", false, false},
		{"E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F", "", true, true},
		{"1234567", "", true, true},
		{"0", "", true, true},
		{"", "", true, true},
	}

	for _, test := range tests {
		if privKey, err := PrivateKeyFromString(test.input); err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error not expected but got: %s", t.Name(), test.input, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error was expected", t.Name(), test.input)
		} else if privKey == nil && !test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and was nil but not expected", t.Name(), test.input)
		} else if privKey != nil && test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and was NOT nil but expected to be nil", t.Name(), test.input)
		} else if privKey != nil && hex.EncodeToString(privKey.Serialize()) != test.expectedKey {
			t.Fatalf("%s Failed: [%s] inputted [%s] expected but failed comparison of keys, got: %s", t.Name(), test.input, test.expectedKey, hex.EncodeToString(privKey.Serialize()))
		}
	}
}

// TestCreatePrivateKey will test the method CreatePrivateKey()
func TestCreatePrivateKey(t *testing.T) {
	rawKey, err := CreatePrivateKey()
	assert.NoError(t, err)
	assert.NotNil(t, rawKey)
	assert.Equal(t, 32, len(rawKey.Serialize()))
}
