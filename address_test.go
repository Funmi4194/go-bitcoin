package bitcoin

import (
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/assert"
)

func TestGetAddressFromScript(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		networkType     NetworkType
		script          string
		expectedAddress string
		expectedError   bool
	}{
		{script: "5120c63ae2b830aee511ff6b3c606b53d6ca0e22d6a7516ad506b5346bf74fa5e3ae", networkType: Mainnet, expectedAddress: "bc1pccaw9wps4mj3rlmt83sxk57keg8z9448294d2p44x34lwna9uwhqe2lxas", expectedError: false},
		{script: "76a91443c1c9de50e52e35546084083363b4586782cbf388ac", networkType: Mainnet, expectedAddress: "17BGRWzKtPstTTvJK9rDusVNMTbWit52XT", expectedError: false},
		{script: "gkckfcbnlkjhoiu7890987654", networkType: Mainnet, expectedAddress: "", expectedError: true},
		{script: "", networkType: Mainnet, expectedAddress: "", expectedError: true},
	}

	for _, test := range tests {
		address, err := GetAddressFromScript(test.script, test.networkType)
		if err != nil && !test.expectedError {
			t.Fatalf("%s Failed: error not expected but got: %s", t.Name(), err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: no error found but error [%v] was expected", t.Name(), test.expectedError)
		} else if address != test.expectedAddress {
			t.Fatalf("%s Failed: expected %s but got %s", t.Name(), test.expectedAddress, address)
		}
	}

}

func TestGetAddressFromPrivateKeyString(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		privKeyHex      string
		networkType     NetworkType
		addressType     AddressType
		expectedAddress string
		expectedError   bool
	}{
		{privKeyHex: "fff9f5137145b7609070fcaf13ab2db3974230699c74b1a3ca5479fb506b5de9", networkType: Mainnet, addressType: NativeSegwit, expectedAddress: "bc1q6n4carl04gn4gr6wslwn42yyfpsvmw0d4w6s7v", expectedError: false},
		{privKeyHex: "fff9f5137145b7609070fcaf13ab2db3974230699c74b1a3ca5479fb506b5de9", networkType: Mainnet, addressType: Taproot, expectedAddress: "bc1psmzc872vcy7l2dcrelmtnw7ujxrt2kuznm9adezhcsaf0x4srwtql74tu6", expectedError: false},
		{privKeyHex: "caf13ab2db3974230699c74b1a3ca5479fb506b5de9", networkType: Mainnet, addressType: NativeSegwit, expectedAddress: "", expectedError: true},
		{privKeyHex: "", networkType: Mainnet, addressType: NativeSegwit, expectedAddress: "", expectedError: true},
	}

	for _, test := range tests {
		address, err := GetAddressFromPrivateKeyString(test.privKeyHex, test.addressType, test.networkType)
		if err != nil && !test.expectedError {
			t.Fatalf("%s Failed: error not expected but got: %s", t.Name(), err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: no error found but error [%v] was expected", t.Name(), test.expectedError)
		} else if address != test.expectedAddress {
			t.Fatalf("%s Failed: expected %s but got %s", t.Name(), test.expectedAddress, address)
		}
	}

}

func TestGetAddressFromMnemonic(t *testing.T) {
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
		address, err := GetAddressFromMnemonic(test.networkType, test.addressType, test.mnemonic, test.password, test.index)
		if err != nil && !test.expectedError {
			t.Fatalf("%s Failed: error not expected but got: %s", t.Name(), err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: no error found but error [%v] was expected", t.Name(), test.expectedError)
		} else if address != test.expectedAddress {
			t.Fatalf("%s Failed: expected %s but got %s", t.Name(), test.expectedAddress, address)
		}
	}
}

// TestGetAddressFromPrivateKeyCompression will test the method GetAddressFromPrivateKey()
func TestGetAddressFromPrivateKey(t *testing.T) {

	privateKey, err := btcec.NewPrivateKey()
	assert.NoError(t, err)

	var address string
	address, err = GetAddressFromPrivateKey(privateKey, Taproot, Mainnet)
	assert.NoError(t, err)
	assert.NotEqual(t, "", address)

	address, err = GetAddressFromPrivateKey(&btcec.PrivateKey{}, Taproot, Mainnet)
	assert.Error(t, err)
	assert.Equal(t, "", address)
}

// TestGetAddressFromPubKeyString will test the method GetAddressFromPubKeyString()
func TestGetAddressFromPubKeyString(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		inputKey        string
		addressType     AddressType
		networkType     NetworkType
		expectedAddress string
		expectedNil     bool
		expectedError   bool
	}{
		{"", Taproot, Mainnet, "", true, true},
		{"0", Segwit, Mainnet, "", true, true},
		{"00000", NativeSegwit, Mainnet, "", true, true},
		{"022d35c7ede60cb68dee6e60ab9ad5863a7a726b297273d7a99b9dfb032a10e3f8", NativeSegwit, Mainnet, "bc1q6n4carl04gn4gr6wslwn42yyfpsvmw0d4w6s7v", false, false},
		{"022d35c7ede60cb68dee6e60ab9ad5863a7a726b297273d7a99b9dfb032a10e3f8", Taproot, Mainnet, "bc1psmzc872vcy7l2dcrelmtnw7ujxrt2kuznm9adezhcsaf0x4srwtql74tu6", false, false},
		{"022d35c7ede60cb68dee6e60ab9ad5863a7a726b297273d7a99b9dfb032a10e3f8", Legacy, Mainnet, "1LQpSSwbXMDiscMy4eamjDFqu5y8DFxcLo", false, false},
	}

	for _, test := range tests {
		if address, err := GetAddressFromPubKeyString(test.inputKey, test.addressType, test.networkType); err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%v] inputted and error not expected but got: %s", t.Name(), test.inputKey, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%v] inputted and error was expected", t.Name(), test.inputKey)
		} else if address != "" && test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and nil was expected", t.Name(), test.inputKey)
		} else if address == "" && !test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and nil was NOT expected", t.Name(), test.inputKey)
		} else if address != "" && address != test.expectedAddress {
			t.Fatalf("%s Failed: [%v] inputted [%s] expected but failed comparison of addresses, got: %s", t.Name(), test.inputKey, test.expectedAddress, address)
		}
	}

}
