package bitcoin

import (
	"strings"
	"testing"
)

func TestGenerateMnemonic(t *testing.T) {
	t.Parallel()

	// this allows multiple test cases
	var tests = []struct {
		bitSize                BitSize
		expectedMnemonicLength int
		expectedError          bool
	}{
		{18, 0, true},
		{12, 0, true},
		{Entropy128, 12, false},
		{Entropy256, 24, false},
	}

	for _, test := range tests {
		mnemonic, err := GenerateMnemonic(test.bitSize)
		if err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%v] inputted and error not expected but got: %s", t.Name(), test.bitSize, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%v] inputted and error was expected", t.Name(), test.bitSize)
		} else if len(strings.Fields(mnemonic)) != test.expectedMnemonicLength {
			t.Fatalf("%s Failed: [%v] inputted and [%v] expected, but got: %v", t.Name(), test.bitSize, test.expectedMnemonicLength, len(strings.Split(mnemonic, " ")))
		}
	}
}
