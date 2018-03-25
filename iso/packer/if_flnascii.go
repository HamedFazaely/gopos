package packer

import (
	"fmt"

	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/util"
)

//FixedLEngthNumberASCIIPacker packs a fixed length number to ascii
type FixedLengthNumberASCIIPacker struct {
}

func (flnap FixedLengthNumberASCIIPacker) Pack(c iso.Component) ([]byte, error) {
	if numeric := util.IsNumericString(c.GetValue()); !numeric {
		return nil, fmt.Errorf(NOT_NUMERIC_STRING_ERROR)
	}
	if len(c.GetValue()) != c.GetMaxLength() {
		return nil, fmt.Errorf(LENGTH_MISSMATCH_ERROR)
	}
	return []byte(c.GetValue()), nil
}
