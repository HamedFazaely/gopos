package util

import (
	"fmt"
	"sort"

	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/ierrors"
)

//GetPackedMTI packs and returns the MTI field
func GetPackedMTI(c iso.Component) ([]byte, error) {

	if !c.IsComposite() {
		return nil, ierrors.ErrLeafCompNoChild
	}

	mti, exists := c.GetComponent(0)

	if !exists {
		return nil, ierrors.ErrNoMTIFieldPresent
	}
	return mti.Pack()
}

//CalculateBitmap calculates the ISOMessage bitmap based on the provided Bitmap interface
func CalculateBitmap(c iso.Component, bm iso.Bitmap) ([]byte, error) {
	if !c.IsComposite() {
		return nil, ierrors.ErrLeafCompNoBitmap
	}
	return bm.Compute(c)

}

//SortAndPackFields sorts the ISOMessage field based on their field number
//and then packs all the fields and returns them in the sorted order
func SortAndPackFields(c iso.Component) ([][]byte, error) {
	if !c.IsComposite() {
		return nil, ierrors.ErrLeafCompNoChild
	}

	var fnos []int //store field numbers to sort them
	var results [][]byte

	for fno := range c.GetChildren() {
		fnos = append(fnos, fno)
	}
	sort.Ints(fnos)
	for _, fno := range fnos {
		if fno == 0 {
			continue
		}
		p, err := c.GetChildren()[fno].Pack()
		if err != nil {
			return nil, err
		}
		results = append(results, p)
	}
	return results, nil
}

//BArrToString generates a bit visible string representation of the byte arrary
//good for logging bitmaps and such
func BArrToString(bytes []byte) string {
	var res string
	for _, b := range bytes {
		res += fmt.Sprintf("%08b ", b)
	}
	return res
}
