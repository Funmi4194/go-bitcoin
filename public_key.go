package bitcoin

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/txscript"
	"github.com/tyler-smith/go-bip39"
)

/*
CreateWalletFromMnemonic creates a new bitcoin wallet address from an addressIndex and mnemonic phrase.
MnemonicPassword can be an empty string if not required
NetworkType can be TESTNET or MAINNET
*/
func CreateWalletFromMnemonic(networkType NetworkType, addressType AddressType, mnemonic, mnemonicPassword string, addressIndex uint32) (string, error) {

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

	return DeriveBitcoinAddressType(childPubKey, addressType, networkType)
}

func DeriveBitcoinAddressType(pubKey *btcec.PublicKey, addressType AddressType, networkType NetworkType) (string, error) {
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
