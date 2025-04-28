package bitcoin

import (
	"errors"
)

// ErrInvalidEntrophy is returned when an invalid mnmonic entrophy bitsize is provided
var ErrInvalidEntrophy = errors.New("invalid entrophy bitsize")

// ErrIncorrectAddressType is returned when an invalid address type iws provided
var ErrIncorrectAddressType = errors.New("incorrect bitcoin address type")

// ErrMissingScript is returned when a script is missing
var ErrMissingScript = errors.New("missing script")

// ErrPrivateKeyMissing is returned when a private key is missing
var ErrPrivateKeyMissing = errors.New("private key is missing")

// ErrMissingPubKey is returned when a pubkey is missing
var ErrMissingPubKey = errors.New("missing pubkey")

// ErrInvalidPubKey is returned when a pubkey is invalid
var ErrInvalidPubKey = errors.New("invalid pubkey")

// ErrMissingAddress is returned when an address is missing
var ErrMissingAddress = errors.New("missing address")

// ErrWifMissing is returned when a wif is missing
var ErrWifMissing = errors.New("wif is missing")
