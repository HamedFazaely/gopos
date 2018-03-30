package packer

import (
	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/ierrors"
	"github.com/HamedFazaely/gopos/iso/logger"
)

//FixedLengthASCII packs a fixed length field into ASCII
type FixedLengthASCII struct {
	packerBase
}

//NewFixedLengthASCII is a factory
func NewFixedLengthASCII(lgr logger.Logger) *FixedLengthASCII {
	return &FixedLengthASCII{
		packerBase: packerBase{
			Logger: lgr,
		},
	}
}

//Pack packs fixed length ascii field
func (fla *FixedLengthASCII) Pack(c iso.Component) ([]byte, error) {

	if ascii := fla.IsASCII(c.GetValue()); !ascii {
		fla.logMessage(ierrors.ErrEncodingError, c)
		return nil, ierrors.ErrEncodingError
	}

	if len(c.GetValue()) != c.GetMaxLength() {
		fla.logMessage(ierrors.ErrLengthMissMatch, c)
		return nil, ierrors.ErrLengthMissMatch
	}
	return []byte(c.GetValue()), nil

}
