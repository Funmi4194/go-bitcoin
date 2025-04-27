package bitcoin

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
)

// GetScriptFromAddress creates an output script from any address type string
func GetScriptFromAddress(address string, networkType NetworkType) (string, error) {

	// invalid address
	if len(address) == 0 {
		return "", ErrMissingAddress
	}

	addr, err := btcutil.DecodeAddress(address, networkType)
	if err != nil {
		return "", err
	}

	// PayToAddrScript returns the standard PkScript for an address
	// It works for all address types
	script, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(script), nil
}
