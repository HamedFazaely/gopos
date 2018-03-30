package ierrors

import "errors"

var (
	ErrLeafCompNoChild     = errors.New("leaf field has no children")
	ErrLeafCompNoBitmap    = errors.New("leaf field can not generate bitmap")
	ErrNoMTIFieldPresent   = errors.New("no MTI field present")
	ErrNoneASCIIValue      = errors.New("only ASCII characters are allowed")
	ErrLengthMissMatch     = errors.New("invalid input length")
	ErrEncodingError       = errors.New("invalid input encoding")
	ErrNotDigit            = errors.New("invalid numeric string")
	ErrPackerNotCompatible = errors.New("packer and component are not compatible")
)
