package packer

import (
	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/ierrors"
	"github.com/HamedFazaely/gopos/iso/logger"
)

//FixedLengthBCDPacker packs an ascii string to bcd
type FixedLengthBCDPacker struct {
	packerBase
}

//NewFixedLengthBCDPacker is a factory
func NewFixedLengthBCDPacker(lgr logger.Logger) *FixedLengthBCDPacker {
	return &FixedLengthBCDPacker{
		packerBase: packerBase{
			Logger: lgr,
		},
	}
}

//Pack packs the component
func (flbp *FixedLengthBCDPacker) Pack(c iso.Component) ([]byte, error) {
	if len(c.GetValue()) != c.GetMaxLength() {
		flbp.logMessage(ierrors.ErrLengthMissMatch, c)
		return nil, ierrors.ErrLengthMissMatch
	}
	res, err := flbp.ToBCD(c.GetValue())
	if err != nil {
		flbp.logMessage(err, c)
		return nil, err
	}
	return res, nil
}
