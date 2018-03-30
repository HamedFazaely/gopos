package misc

import (
	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/ierrors"
	"github.com/HamedFazaely/gopos/iso/logger"
	"github.com/HamedFazaely/gopos/iso/util"
)

//ISOBitmap implements Bitmap interface
type ISOBitmap struct {
	logger.Logger
}

//NewBitMap is bitmap factory
func NewBitMap(lgr logger.Logger) *ISOBitmap {
	return &ISOBitmap{
		Logger: lgr,
	}
}

//Compute simply computes the bitmap
func (b *ISOBitmap) Compute(c iso.Component) ([]byte, error) {
	max, err := getMaxField(c)
	if err != nil {
		b.Log(logger.Info|logger.Trace, err.Error())
		return nil, err
	}
	if max > 128 {
		b.Log(logger.Trace|logger.Warning, ierrors.TooLongBitmap)
		panic(ierrors.TooLongBitmap)
	}
	bml := 8
	if max > 64 {
		bml *= 2
	}
	arr := make([]byte, bml)
	if max > 64 {
		arr[0] = 0x80 // 10000000 setting field no 1 to indicate there is a secondary bitmap
	}
	for fno := range c.GetChildren() {
		setBit(fno, arr)
	}
	b.Log(logger.Trace|logger.Info, "bitmap generated %s", util.BArrToString(arr))
	return arr, nil

}

func setBit(fno int, bm []byte) {
	if fno > 0 {
		idx := fno / 8 // index of the byte containing the bit
		r := fno % 8
		bitPos := uint8(r - 1)
		if r == 0 {
			bitPos = uint8(7)
			idx--
		}
		bm[idx] |= byte(0x80 >> bitPos)
	}

}

func getMaxField(c iso.Component) (int, error) {
	if !c.IsComposite() {
		return 0, ierrors.ErrLeafCompNoChild
	}
	max := 0
	for fno := range c.GetChildren() {
		if fno > max {
			max = fno
		}
	}
	return max, nil
}
