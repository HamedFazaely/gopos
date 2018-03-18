package misc

import (
	"fmt"
	"gopos/iso"
)

type ISOBitmap struct {
}

func (b *ISOBitmap) Compute(c iso.Component) ([]byte, error) {
	max, err := getMaxField(c)
	if err != nil {
		return nil, err
	}
	if max > 128 {
		panic("only 128 length bitmap is supported")
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
		bm[idx] ^= byte(0x80 >> bitPos)
	}

}

func getMaxField(c iso.Component) (int, error) {
	if !c.IsComposite() {
		return 0, fmt.Errorf("leaf fields do not have any children")
	}
	fnos := make([]int, len(c.GetChildren()))
	for n := range c.GetChildren() {
		fnos = append(fnos, n)
	}
	return fnos[len(fnos)-1], nil
}
