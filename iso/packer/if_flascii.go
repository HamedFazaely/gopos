package packer

import (
	"fmt"

	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/util"
)

//FixedLengthASCII packs a fixed length field into ASCII
type FixedLengthASCII struct {
}

func (fla FixedLengthASCII) Pack(c iso.Component) ([]byte, error) {

	if ascii := util.IsASCII(c.GetValue()); !ascii {
		return nil, fmt.Errorf(ENCODING_ERROR)
	}

	if len(c.GetValue()) != c.GetMaxLength() {
		return nil, fmt.Errorf(LENGTH_MISSMATCH_ERROR)
	}
	return []byte(c.GetValue()), nil

}
