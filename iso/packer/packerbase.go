package packer

import (
	"strconv"

	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/ierrors"
	"github.com/HamedFazaely/gopos/iso/logger"
	"github.com/HamedFazaely/gopos/iso/util"
)

type packerBase struct {
	logger.Logger
}

func (pb *packerBase) IsNumericString(s string) bool {
	return util.IsNumericString(s)
}

func (pb *packerBase) IsASCII(s string) bool {
	return util.IsASCII(s)
}

func (pb *packerBase) ToBCD(num string) ([]byte, error) {
	return util.ASCIIToBCD(num)
}

func (pb *packerBase) logMessage(err error, c iso.Component) {
	pb.Log(logger.Trace|logger.Info|logger.Error, "%s fno : %d value : %s maxlen : %d", err.Error(), c.GetKey(), c.GetValue(), c.GetMaxLength())
}

func (pb *packerBase) logError(err error) {
	pb.Log(logger.Error, err.Error())
}

func (pb *packerBase) GetPackedMTI(c iso.Component) ([]byte, error) {
	return util.GetPackedMTI(c)
}

func (pb *packerBase) numToString(n int) string {
	return strconv.FormatInt(int64(n), 10)
}

type LengthFormatter interface {
	Format(s string) ([]byte, error)
}

type NumASCIIFormatter struct {
}

func (f *NumASCIIFormatter) Format(s string) ([]byte, error) {
	if numeric := util.IsNumericString(s); !numeric {
		return nil, ierrors.ErrNotDigit
	}
	return []byte(s), nil
}

type NumBCDFormatter struct {
}

func (f *NumBCDFormatter) Format(s string) ([]byte, error) {
	return util.ASCIIToBCD(s)
}
