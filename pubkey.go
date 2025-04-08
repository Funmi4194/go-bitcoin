package bitcoin

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec/v2"
)

// PubKeyFromString will convert a pubKey (hex string) into a pubkey (*bec.PublicKey)
func PubKeyFromString(pubKeyHex string) (*btcec.PublicKey, error) {

	// invalid pubKey
	if len(pubKeyHex) == 0 {
		return nil, ErrMissingPubKey
	}

	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return nil, err
	}

	// create a public key from bytes
	pubKey, err := btcec.ParsePubKey(pubKeyBytes)
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}

// PubKeyFromPrivateKeyString will derive a pubKey (hex encoded) from a given private key
func PubKeyFromPrivateKeyString(privateKey string) (string, error) {
	rawKey, err := PrivateKeyFromString(privateKey)
	if err != nil {
		return "", err
	}

	return PubKeyFromPrivateKey(rawKey), nil
}

// PubKeyFromPrivateKey will derive a pubKey (hex encoded) from a given private key
func PubKeyFromPrivateKey(privateKey *btcec.PrivateKey) string {
	return hex.EncodeToString(privateKey.PubKey().SerializeCompressed())
}
