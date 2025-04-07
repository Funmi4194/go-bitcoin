package bitcoin

import "github.com/btcsuite/btcd/chaincfg"

// AddressType denotes the different bitcoin
type AddressType string

const (
	Legacy       AddressType = "P2PKH"
	Segwit       AddressType = "P2SH"
	NativeSegwit AddressType = "P2WPKH"
	Taproot      AddressType = "P2TR"
)

// NetworkType wraps chaincfg.Params to allow type safety in functions
type NetworkType *chaincfg.Params

var (
	Mainnet NetworkType = &chaincfg.MainNetParams
	Testnet NetworkType = &chaincfg.TestNet3Params
)

type BitSize int

const (
	Entropy128 BitSize = 128 //12-word mnemonic
	Entropy256 BitSize = 256 //24-word mnemonic
)
