package packer

import (
	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/ierrors"
	"github.com/HamedFazaely/gopos/iso/logger"
)

//FixedLengthNumberASCIIPacker packs a fixed length number to ascii
type FixedLengthNumberASCIIPacker struct {
	packerBase
}

//NewFixedLengthNumberASCIIPacker is a factory
func NewFixedLengthNumberASCIIPacker(lgr logger.Logger) *FixedLengthNumberASCIIPacker {
	return &FixedLengthNumberASCIIPacker{
		packerBase: packerBase{
			Logger: lgr,
		},
	}
}

//Pack packs the field
func (flnap *FixedLengthNumberASCIIPacker) Pack(c iso.Component) ([]byte, error) {
	if numeric := flnap.IsNumericString(c.GetValue()); !numeric {
		flnap.logMessage(ierrors.ErrNotDigit, c)
		return nil, ierrors.ErrNotDigit
	}
	if len(c.GetValue()) != c.GetMaxLength() {
		flnap.logMessage(ierrors.ErrLengthMissMatch, c)
		return nil, ierrors.ErrLengthMissMatch
	}
	return []byte(c.GetValue()), nil
}
