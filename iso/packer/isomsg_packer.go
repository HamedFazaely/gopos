package packer

import (
	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/misc"
	"github.com/HamedFazaely/gopos/iso/util"
)

type ISOMessagePacker struct {
}

func (im ISOMessagePacker) Pack(c iso.Component) ([]byte, error) {

	if !c.IsComposite() {
		panic("this packer must be used with outer most iso message")
	}
	result, err := util.GetPackedMTI(c)
	if err != nil {
		panic("mti field is not present")
	}
	var mapper misc.ISOBitmap
	bm, err := util.CalculateBitmap(c, &mapper)
	if err != nil {
		return nil, err
	}
	result = append(result, bm...)
	sf, err := util.SortAndPackFields(c)
	if err != nil {
		return nil, err
	}
	for _, p := range sf {
		result = append(result, p...)
	}
	return result, nil
}
