package bitcoin

import (
	"errors"
)

// ErrInvalidEntrophy is returned when an invalid mnmonic entrophy bitsize is provided
var ErrInvalidEntrophy = errors.New("invalid entrophy bitsize")

// ErrIncorrectAddressType is returned when an invalid address type iws provided
var ErrIncorrectAddressType = errors.New("incorrect bitcoin address type")
