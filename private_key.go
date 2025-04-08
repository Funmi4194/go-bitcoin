package bitcoin

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec/v2"
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
