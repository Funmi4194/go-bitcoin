package bitcoin

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

// TestCreatePrivateKeyString will test the method CreatePrivateKeyString()
func TestCreatePrivateKeyString(t *testing.T) {
	key, err := CreatePrivateKeyString()
	assert.NoError(t, err)
	assert.Equal(t, 64, len(key))
}

func TestCreateWif(t *testing.T) {
	t.Run("TestCreateWif", func(t *testing.T) {
		t.Parallel()

		// create a WIF
		wifKey, err := CreateWif(Mainnet)
		require.NoError(t, err)
		require.NotNil(t, wifKey)

		require.Equalf(t, 51, len(wifKey.String()), "WIF should be 51 characters long but got: %d", len(wifKey.String()))
	})

	t.Run("TestWifToPrivateKey", func(t *testing.T) {
		t.Parallel()

		// create a WIF
		wifKey, err := CreateWif(Mainnet)
		require.NoError(t, err)
		require.NotNil(t, wifKey)
		require.Equalf(t, 51, len(wifKey.String()), "WIF should be 51 characters long but got: %d", len(wifKey.String()))

		// convert WIF to private key
		var privateKey *btcec.PrivateKey
		privateKey, err = WifToPrivateKey(wifKey.String())
		require.NoError(t, err)
		require.NotNil(t, privateKey)
		privateKeyString := hex.EncodeToString(privateKey.Serialize())
		require.Equalf(t, 64, len(privateKeyString), "Private Key should be 64 characters long, got: %d", len(privateKeyString))
	})
}

// TestCreateWifString will test the method CreateWifString()
func TestCreateWifString(t *testing.T) {
	t.Run("TestCreateWifString", func(t *testing.T) {
		t.Parallel()

		// Create a WIF
		wifKey, err := CreateWifString(Mainnet)
		require.NoError(t, err)
		require.NotNil(t, wifKey)
		// t.Log("WIF:", wifKey)
		require.Equalf(t, 51, len(wifKey), "WIF should be 51 characters long, got: %d", len(wifKey))
	})

	t.Run("TestWifToPrivateKeyString", func(t *testing.T) {
		t.Parallel()

		// Create a WIF
		wifKey, err := CreateWifString(Mainnet)
		require.NoError(t, err)
		require.NotNil(t, wifKey)
		// t.Log("WIF:", wifKey)
		require.Equalf(t, 51, len(wifKey), "WIF should be 51 characters long, got: %d", len(wifKey))

		// Convert WIF to Private Key
		var privateKeyString string
		privateKeyString, err = WifToPrivateKeyString(wifKey)
		require.NoError(t, err)
		require.NotNil(t, privateKeyString)
		// t.Log("Private Key:", privateKeyString)
		require.Equalf(t, 64, len(privateKeyString), "Private Key should be 64 characters long, got: %d", len(privateKeyString))

	})
}

// TestPrivateAndPublicKeys will test the method PrivateAndPublicKeys()
func TestGetPrivateAndPublicKeys(t *testing.T) {

	t.Parallel()

	var tests = []struct {
		privateKey         string
		expectedPrivateKey string
		expectedNil        bool
		expectedError      bool
	}{
		{"", "", true, true},
		{"0", "", true, true},
		{"00000", "", true, true},
		{"0-0-0-0-0", "", true, true},
		{"z4035dd4c7dda99ac473905a3d82f7864322b49bab1ff441cc457183b9bd8abz", "", true, true},
		{"8e4588bce9a1b4d9d518fbab91bf1d07dfd7805a4f7c0acf35d3476587af03af", "8e4588bce9a1b4d9d518fbab91bf1d07dfd7805a4f7c0acf35d3476587af03af", false, false},
	}

	for _, test := range tests {
		if privateKey, publicKey, err := GetPrivateAndPublicKeys(test.privateKey); err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error not expected but got: %s", t.Name(), test.privateKey, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error was expected", t.Name(), test.privateKey)
		} else if (privateKey == nil || publicKey == nil) && !test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and was nil but not expected", t.Name(), test.privateKey)
		} else if (privateKey != nil || publicKey != nil) && test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and was NOT nil but expected to be nil", t.Name(), test.privateKey)
		} else if privateKey != nil && hex.EncodeToString(privateKey.Serialize()) != test.expectedPrivateKey {
			t.Fatalf("%s Failed: [%s] inputted [%s] expected but failed comparison of keys, got: %s", t.Name(), test.privateKey, test.expectedPrivateKey, hex.EncodeToString(privateKey.Serialize()))
		}
	}
}

// TestPrivateKeyToWif will test the method PrivateKeyToWif()
func TestPrivateKeyToWif(t *testing.T) {

	t.Parallel()

	var tests = []struct {
		privateKey    string
		expectedWif   string
		expectedNil   bool
		expectedError bool
	}{
		{"", "", true, true},
		{"0", "", true, true},
		{"000000", "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAbuatmU", false, false},
		{"6D792070726976617465206B6579", "5HpHagT65TZzG1PH3CSu63k8DbuTZnNJf6HgyQNymvXmALAsm9s", false, false},
		{"54035dd4c7dda99ac473905a3d82f7864322b49bab1ff441cc457183b9bd8azz", "", true, true},
		{"482614ee78fc5c78e0c2c8394da20a517155a016f229913aa27de5071993698f", "5JN4XPF6P6CZWZJrBHEeLTHtRZmAU8TZPyXqcckmaE5NzfmNLA6", false, false},
	}

	for _, test := range tests {
		if privateWif, err := PrivateKeyToWif(test.privateKey, Mainnet); err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error not expected but got: %s", t.Name(), test.privateKey, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error was expected", t.Name(), test.privateKey)
		} else if privateWif == nil && !test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and was nil but not expected", t.Name(), test.privateKey)
		} else if privateWif != nil && test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and was NOT nil but expected to be nil", t.Name(), test.privateKey)
		} else if privateWif != nil && privateWif.String() != test.expectedWif {
			t.Fatalf("%s Failed: [%s] inputted [%s] expected but failed comparison of keys, got: %s", t.Name(), test.privateKey, test.expectedWif, privateWif.String())
		}
	}
}

// TestPrivateKeyToWifString will test the method PrivateKeyToWifString()
func TestPrivateKeyToWifString(t *testing.T) {

	t.Parallel()

	var tests = []struct {
		privateKey    string
		expectedWif   string
		expectedError bool
	}{
		{"", "", true},
		{"0", "", true},
		{"1d3f52ebd13d6879b96c6e03ba69b9fad91294d19f5f7ca3184bbc8d50ced2c7", "5J3AfapLfHyj2N4drvx9Dhafondb59fCPB3HBUvVixTKy2YzR2v", false},
		{"81641e606078bd9ecc08f199182474699366d9b2c21e83cc12fb04ce20c7e9e1", "5JoGhjqC3SMdAqTwu7AHfx23rhESZmK8VTwEhfT98Dd79iJtT5X", false},
		{"54035dd4c7dda99ac473905a3d82f7864322b49bab1ff441cc457183b9bd8azz", "", true},
		{"482614ee78fc5c78e0c2c8394da20a517155a016f229913aa27de5071993698f", "5JN4XPF6P6CZWZJrBHEeLTHtRZmAU8TZPyXqcckmaE5NzfmNLA6", false},
	}

	for _, test := range tests {
		if privateWif, err := PrivateKeyToWifString(test.privateKey, Mainnet); err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error not expected but got: %s", t.Name(), test.privateKey, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error was expected", t.Name(), test.privateKey)
		} else if privateWif != test.expectedWif {
			t.Fatalf("%s Failed: [%s] inputted [%s] expected but failed comparison of keys, got: %s", t.Name(), test.privateKey, test.expectedWif, privateWif)
		}
	}
}

// TestWifToPrivateKey will test the method WifToPrivateKey()
func TestWifToPrivateKey(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		WifKey        string
		expectedKey   string
		expectedNil   bool
		expectedError bool
	}{
		{"", "", true, true},
		{"0", "", true, true},
		{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAbuatmU", "0000000000000000000000000000000000000000000000000000000000000000", false, false},
		{"5HpHagT65TZzG1PH3CSu63k8DbuTZnNJf6HgyQNymvXmALAsm9s", "0000000000000000000000000000000000006d792070726976617465206b6579", false, false},
		{"54035dd4c7dda99ac473905a3d82f7864322b49bab1ff441cc457183b9bd8azz", "", true, true},
		{"5JTHas7yTFMBLqgFogxZFf8Vc5uKEbkE7yQAQ2g3xPHo2sNG1Ei", "54035dd4c7dda99ac473905a3d82f7864322b49bab1ff441cc457183b9bd8abd", false, false},
	}

	for _, test := range tests {
		if privateKey, err := WifToPrivateKey(test.WifKey); err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error not expected but got: %s", t.Name(), test.WifKey, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error was expected", t.Name(), test.WifKey)
		} else if privateKey == nil && !test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and was nil but not expected", t.Name(), test.WifKey)
		} else if privateKey != nil && test.expectedNil {
			t.Fatalf("%s Failed: [%s] inputted and was NOT nil but expected to be nil", t.Name(), test.WifKey)
		} else if privateKey != nil && hex.EncodeToString(privateKey.Serialize()) != test.expectedKey {
			t.Fatalf("%s Failed: [%s] inputted [%s] expected but failed comparison of keys, got: %s", t.Name(), test.WifKey, test.expectedKey, hex.EncodeToString(privateKey.Serialize()))
		}
	}
}

// TestWifToPrivateKeyString will test the method WifToPrivateKeyString()
func TestWifToPrivateKeyString(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		wifKey        string
		expectedKey   string
		expectedError bool
	}{
		{"", "", true},
		{"0", "", true},
		{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAbuatmU", "0000000000000000000000000000000000000000000000000000000000000000", false},
		{"5HpHagT65TZzG1PH3CSu63k8DbuTZnNJf6HgyQNymvXmALAsm9s", "0000000000000000000000000000000000006d792070726976617465206b6579", false},
		{"54035dd4c7dda99ac473905a3d82f7864322b49bab1ff441cc457183b9bd8azz", "", true},
		{"5JTHas7yTFMBLqgFogxZFf8Vc5uKEbkE7yQAQ2g3xPHo2sNG1Ei", "54035dd4c7dda99ac473905a3d82f7864322b49bab1ff441cc457183b9bd8abd", false},
	}

	for _, test := range tests {
		if privateKey, err := WifToPrivateKeyString(test.wifKey); err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error not expected but got: %s", t.Name(), test.wifKey, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%s] inputted and error was expected", t.Name(), test.wifKey)
		} else if privateKey != test.expectedKey {
			t.Fatalf("%s Failed: [%s] inputted [%s] expected but failed comparison of keys, got: %s", t.Name(), test.wifKey, test.expectedKey, privateKey)
		}
	}
}

// TestWifFromString will test the method WifFromString()
func TestWifFromString(t *testing.T) {
	t.Run("TestCreateWifFromPrivateKey", func(t *testing.T) {
		t.Parallel()

		// Create a Private Key
		privateKey, err := CreatePrivateKeyString()
		require.NoError(t, err)
		require.NotNil(t, privateKey)

		// Create a WIF
		var wifKey *btcutil.WIF
		wifKey, err = PrivateKeyToWif(privateKey, Mainnet)
		require.NoError(t, err)
		require.NotNil(t, wifKey)
		wifKeyString := wifKey.String()
		t.Log("WIF:", wifKeyString)
		require.Equalf(t, 51, len(wifKeyString), "WIF should be 51 characters long, got: %d", len(wifKeyString))

		// Convert WIF to Private Key
		var privateKeyString string
		privateKeyString, err = WifToPrivateKeyString(wifKeyString)
		require.NoError(t, err)
		require.NotNil(t, privateKeyString)
		t.Log("Private Key:", privateKeyString)
		require.Equalf(t, 64, len(privateKeyString), "Private Key should be 64 characters long, got: %d", len(privateKeyString))

		// Compare Private Keys
		require.Equalf(t, privateKey, privateKeyString, "Private Key should be equal, got: %s", privateKeyString)

		// Decode WIF
		var decodedWif *btcutil.WIF
		decodedWif, err = WifFromString(wifKeyString)
		require.NoError(t, err)
		require.NotNil(t, decodedWif)
		require.Equalf(t, wifKeyString, decodedWif.String(), "WIF should be equal, got: %s", decodedWif.String())
	})

	t.Run("TestWifFromStringMissingWIF", func(t *testing.T) {
		t.Parallel()

		_, err := WifFromString("")
		require.Error(t, err)
		require.Equal(t, ErrWifMissing, err)
	})

	t.Run("TestWifFromStringInvalidWIF", func(t *testing.T) {
		t.Parallel()

		_, err := WifFromString("invalid")
		require.Error(t, err)
		require.Equal(t, "malformed private key", err.Error())
	})
}
