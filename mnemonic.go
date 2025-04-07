package bitcoin

import "github.com/tyler-smith/go-bip39"

// NewMnemonicFromEntropy returns a BIP-39 mnemonic from entropy.
func NewMnemonicFromEntropy(entropy []byte) (string, error) {
	return bip39.NewMnemonic(entropy)
}

func GenerateMnemonic(bitSize BitSize) (string, error) {

	if bitSize != Entropy128 && bitSize != Entropy256 {
		return "", ErrInvalidEntrophy
	}

	// generate a random entropy of the provide bit size (e.g 256 provides 24 worda)
	entropy, err := bip39.NewEntropy(int(bitSize))
	if err != nil {
		return "", err
	}

	mnemonic, err := NewMnemonicFromEntropy(entropy)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}
