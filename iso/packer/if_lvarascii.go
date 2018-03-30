package packer

import (
	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/ierrors"
	"github.com/HamedFazaely/gopos/iso/logger"
)

type LVARASCIIStringPacker struct {
	packerBase
	lfmt LengthFormatter
}

//NewLVARASCIIStringPacker is a factory
func NewLVARASCIIStringPacker(lgr logger.Logger, lf LengthFormatter) *LVARASCIIStringPacker {
	return &LVARASCIIStringPacker{
		packerBase: packerBase{
			Logger: lgr,
		},
		lfmt: lf,
	}
}

func (lvasp *LVARASCIIStringPacker) Pack(c iso.Component) ([]byte, error) {

	if ascii := lvasp.IsASCII(c.GetValue()); !ascii {
		lvasp.logMessage(ierrors.ErrNoneASCIIValue, c)
		return nil, ierrors.ErrNoneASCIIValue
	}

	if c.GetMaxLength() > 9 {
		lvasp.logMessage(ierrors.ErrPackerNotCompatible, c)
		return nil, ierrors.ErrPackerNotCompatible
	}

	if len(c.GetValue()) > c.GetMaxLength() {
		lvasp.logMessage(ierrors.ErrLengthMissMatch, c)
		return nil, ierrors.ErrLengthMissMatch
	}

	ns := lvasp.numToString(len(c.GetValue()))
	l, err := lvasp.lfmt.Format(ns)
	if err != nil {
		lvasp.logMessage(err, c)
		return nil, err
	}

	l = append(l, []byte(c.GetValue())...)
	return l, nil

}
