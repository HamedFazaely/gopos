package packer

import (
	"fmt"
	"strconv"

	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/util"
)

type LVARASCIIStringPacker struct {
	LengthFormat LenFormat
}

func (lvasp *LVARASCIIStringPacker) Pack(c iso.Component) ([]byte, error) {
	if ascii := util.IsASCII(c.GetValue()); !ascii {
		return nil, fmt.Errorf(ENCODING_ERROR)
	}
	if c.GetMaxLength() > 9 || len(c.GetValue()) > c.GetMaxLength() {
		return nil, fmt.Errorf(LENGTH_MISSMATCH_ERROR)
	}
	var l []byte
	ns := strconv.FormatInt(int64(len(c.GetValue())), 10)
	switch lvasp.LengthFormat {
	case ASCII:
		l = []byte(ns)
		break
	case BCD:
		r, err := util.ASCIIToBCD(ns)
		if err != nil {
			return nil, err
		}
		l = r
		break
	}

	l = append(l, []byte(c.GetValue())...)
	return l, nil

}
