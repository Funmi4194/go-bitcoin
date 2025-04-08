package bitcoin

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/txscript"
	"github.com/tyler-smith/go-bip39"
)

// GetAddressFromPrivateKey takes a bec private key and returns a Bitcoin address of the required type
func GetAddressFromPrivateKey(privateKey *btcec.PrivateKey, addressType AddressType, network NetworkType) (string, error) {
	return GetAddressFromPubKey(privateKey.PubKey(), addressType, network)
}

// GetAddressFromPrivateKeyString takes a private key string and returns a Bitcoin address of the required type
func GetAddressFromPrivateKeyString(privateKey string, addressType AddressType, network NetworkType) (string, error) {
	rawKey, err := PrivateKeyFromString(privateKey)
	if err != nil {
		return "", err
	}

	return GetAddressFromPubKey(rawKey.PubKey(), addressType, network)
}

// GetAddressFromPubKeyString is a convenience function to use a hex string pubKey
func GetAddressFromPubKeyString(pubKey string, addressType AddressType, network NetworkType) (string, error) {
	rawPubKey, err := PubKeyFromString(pubKey)
	if err != nil {
		return "", err
	}

	return GetAddressFromPubKey(rawPubKey, addressType, network)
}

// GetAddressFromScript will take an output script and extract a standard bitcoin address
func GetAddressFromScript(script string, networkType NetworkType) (string, error) {

	// No script?
	if len(script) == 0 {
		return "", ErrMissingScript
	}

	// Decode the hex string into bytes
	scriptBytes, err := hex.DecodeString(script)
	if err != nil {
		return "", err
	}

	// Extract the addresses from the script
	_, addresses, _, err := txscript.ExtractPkScriptAddrs(scriptBytes, networkType)
	if err != nil {
		return "", err
	}

	// Safety check: Ensure at least one address was extracted.
	// This case is unlikely because the error above should already cover it,
	// but we include this check to avoid using an empty address list below.
	if len(addresses) == 0 {
		return "", fmt.Errorf("invalid output script, missing an address")
	}

	// Use the encoded version of the address
	return addresses[0].EncodeAddress(), nil
}

/*
CreateAddressFromMnemonic creates a new bitcoin wallet address from an addressIndex and mnemonic phrase.
MnemonicPassword can be an empty string if not required
NetworkType can be TESTNET or MAINNET
*/
func GetAddressFromMnemonic(networkType NetworkType, addressType AddressType, mnemonic, mnemonicPassword string, addressIndex uint32) (string, error) {

	// generate a Bip32 HD wallet for the mnemonic
	masterKey, err := hdkeychain.NewMaster(bip39.NewSeed(mnemonic, mnemonicPassword), networkType)
	if err != nil {
		return "", err
	}

	// derive the first child key
	childKey, err := masterKey.Derive(hdkeychain.HardenedKeyStart + addressIndex)
	if err != nil {
		return "", err
	}

	// convert child key to bitcoin address
	childPubKey, err := childKey.ECPubKey()
	if err != nil {
		return "", err
	}

	return GetAddressFromPubKey(childPubKey, addressType, networkType)
}

func GetAddressFromPubKey(pubKey *btcec.PublicKey, addressType AddressType, networkType NetworkType) (string, error) {

	valid := IsValidPublicKey(pubKey)
	if !valid {
		return "", ErrInvalidPubKey
	}

	switch addressType {
	case Legacy:
		// P2PKH (Legacy)
		pubKeyHash := btcutil.Hash160(pubKey.SerializeCompressed())
		addr, err := btcutil.NewAddressPubKeyHash(pubKeyHash, networkType)
		if err != nil {
			return "", err
		}
		return addr.EncodeAddress(), nil
	case Segwit:
		// P2SH (Wrapped SegWit)
		pubKeyHash := btcutil.Hash160(pubKey.SerializeCompressed())
		witnessProg := append([]byte{0x00, 0x14}, pubKeyHash...) // P2WPKH in P2SH
		addr, err := btcutil.NewAddressScriptHash(witnessProg, networkType)
		if err != nil {
			return "", err
		}
		return addr.EncodeAddress(), nil
	case NativeSegwit:
		// P2WPKH (Native SegWit, BECH32)
		pubKeyHash := btcutil.Hash160(pubKey.SerializeCompressed())
		addr, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, networkType)
		if err != nil {
			return "", err
		}
		return addr.EncodeAddress(), nil
	case Taproot:
		// P2TR (Taproot)
		taprootKey := txscript.ComputeTaprootKeyNoScript(pubKey)

		// Create a Taproot address
		addr, err := btcutil.NewAddressTaproot(taprootKey.SerializeCompressed()[1:], networkType)
		if err != nil {
			return "", err
		}

		return addr.EncodeAddress(), nil
	default:
		return "", ErrIncorrectAddressType
	}
}

func IsValidPublicKey(pubKey *btcec.PublicKey) bool {
	if pubKey == nil {
		return false
	}
	return btcec.S256().IsOnCurve(pubKey.X(), pubKey.Y())
}
