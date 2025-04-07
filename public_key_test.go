package bitcoin

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
)

func TestDeriveBitcoinAddressType(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		networkType     NetworkType
		addressType     AddressType
		mnemonic        string
		password        string
		index           uint32
		expectedAddress string
		expectedError   bool
	}{
		{networkType: Mainnet, addressType: Legacy, mnemonic: TestMnemonicPhrase, password: "", index: uint32(1), expectedAddress: "1GvoWk6UPuogoV6SiEY176FrQjHxro58bX", expectedError: false},
		{networkType: Mainnet, addressType: NativeSegwit, mnemonic: TestMnemonicPhrase, password: "", index: uint32(1), expectedAddress: "bc1q46mvy7wpsg0fj5npsc977fk9pcxjyaqawnv37r", expectedError: false},
		{networkType: Mainnet, addressType: Segwit, mnemonic: TestMnemonicPhrase, password: "", index: uint32(1), expectedAddress: "3DhNDPHT2jjrTD6645SHM5YGiMdfwr9Ytt", expectedError: false},
		{networkType: Mainnet, addressType: Taproot, mnemonic: TestMnemonicPhrase, password: "", index: uint32(1), expectedAddress: "bc1pwvlp7wdnxuhh623vc6yy5lu84n5yyc3ymaylmjs3e6mrggw5lteswgvgng", expectedError: false},
		{networkType: Mainnet, addressType: "", mnemonic: TestMnemonicPhrase, password: "", index: uint32(1), expectedAddress: "", expectedError: true},
		{networkType: &chaincfg.Params{}, addressType: Taproot, mnemonic: TestMnemonicPhrase, password: "", index: uint32(1), expectedAddress: "", expectedError: false},
	}

	for _, test := range tests {
		address, err := CreateWalletFromMnemonic(test.networkType, test.addressType, test.mnemonic, test.password, test.index)
		if err != nil && !test.expectedError {
			t.Fatalf("%s Failed: error not expected but got: %s", t.Name(), err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: no error found but error [%v] was expected", t.Name(), test.expectedError)
		} else if address != test.expectedAddress {
			t.Fatalf("%s Failed: expected %s but got %s", t.Name(), test.expectedAddress, address)
		}
	}
}
