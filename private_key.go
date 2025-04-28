package bitcoin

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
)

// PrivateKeyFromString turns a private key (hex encoded string) into an btcec.PrivateKey
func PrivateKeyFromString(privateKey string) (*btcec.PrivateKey, error) {
	if len(privateKey) == 0 {
		return nil, ErrPrivateKeyMissing
	}

	privKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	privKey, _ := btcec.PrivKeyFromBytes(privKeyBytes)
	return privKey, nil
}

// CreatePrivateKey will create a new private key (*btcec.PrivateKey)
func CreatePrivateKey() (*btcec.PrivateKey, error) {
	return btcec.NewPrivateKey()
}

// CreatePrivateKeyString will create a new private key (hex encoded)
func CreatePrivateKeyString() (string, error) {
	privateKey, err := CreatePrivateKey()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(privateKey.Serialize()), nil
}

// CreateWif creates a new WIF (*btcutil.WIF)
func CreateWif(networkType NetworkType) (*btcutil.WIF, error) {
	privateKey, err := CreatePrivateKey()
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(privateKey, networkType, false)
}

// CreateWifString will create a new WIF (string)
func CreateWifString(networkType NetworkType) (string, error) {
	wifKey, err := CreateWif(networkType)
	if err != nil {
		return "", err
	}

	return wifKey.String(), nil
}

// PrivateAndPublicKeys will return both the private and public key in one method
// Expects a hex encoded privateKey
func GetPrivateAndPublicKeys(privateKey string) (*btcec.PrivateKey, *btcec.PublicKey, error) {

	// Missing private key
	if len(privateKey) == 0 {
		return nil, nil, ErrPrivateKeyMissing
	}

	// Decode the private key
	privKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, nil, err
	}

	privKey, pubKey := btcec.PrivKeyFromBytes(privKeyBytes)
	return privKey, pubKey, nil
}

// PrivateKeyToWif will convert a private key to a WIF (*btcutil.WIF)
func PrivateKeyToWif(privateKey string, networkType NetworkType) (*btcutil.WIF, error) {

	// Missing private key
	if len(privateKey) == 0 {
		return nil, ErrPrivateKeyMissing
	}

	// Decode the private key
	privKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	privKey, _ := btcec.PrivKeyFromBytes(privKeyBytes)

	// Create a new WIF (error never gets hit since (networkType) is set correctly)
	return btcutil.NewWIF(privKey, networkType, false)
}

// PrivateKeyToWifString will convert a private key to a WIF (string)
func PrivateKeyToWifString(privateKey string, networkType NetworkType) (string, error) {
	privateWif, err := PrivateKeyToWif(privateKey, networkType)
	if err != nil {
		return "", err
	}

	return privateWif.String(), nil
}

// WifToPrivateKey will convert a WIF to a private key (*bec.PrivateKey)
func WifToPrivateKey(wifKey string) (*btcec.PrivateKey, error) {

	// Missing wif
	if len(wifKey) == 0 {
		return nil, ErrWifMissing
	}

	// Decode the wif
	decodedWif, err := btcutil.DecodeWIF(wifKey)
	if err != nil {
		return nil, err
	}

	// Return the private key
	return decodedWif.PrivKey, nil
}

// WifToPrivateKeyString will convert a WIF to private key (string)
func WifToPrivateKeyString(wif string) (string, error) {

	// Convert the wif to private key
	privateKey, err := WifToPrivateKey(wif)
	if err != nil {
		return "", err
	}

	// Return the hex (string) version of the private key
	return hex.EncodeToString(privateKey.Serialize()), nil
}

// WifFromString will convert a WIF (string) to a WIF (*btcutil.WIF)
func WifFromString(wifKey string) (*btcutil.WIF, error) {

	// Missing wif
	if len(wifKey) == 0 {
		return nil, ErrWifMissing
	}

	// Decode the WIF
	decodedWif, err := btcutil.DecodeWIF(wifKey)
	if err != nil {
		return nil, err
	}

	return decodedWif, nil
}
