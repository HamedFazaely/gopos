package packer

import (
	"fmt"

	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/util"
)

//FixedLengthBCDPacker packs an ascii string to bcd
type FixedLengthBCDPacker struct {
}

func (flbp FixedLengthBCDPacker) Pack(c iso.Component) ([]byte, error) {
	if len(c.GetValue()) != c.GetMaxLength() {
		return nil, fmt.Errorf(LENGTH_MISSMATCH_ERROR)
	}
	res, err := util.ASCIIToBCD(c.GetValue())
	if err != nil {
		return nil, err
	}
	return res, nil
}
