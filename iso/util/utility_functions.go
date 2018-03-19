package util

import (
	"fmt"
	"github.com/HamedFazaely/gopos/iso"
	"sort"
)

func GetPackedMTI(c iso.Component) ([]byte, error) {

	if !c.IsComposite() {
		return nil, fmt.Errorf("leaf fields do not have any children")
	}

	mti, exists := c.GetComponent(0)

	if !exists {
		return nil, fmt.Errorf("mti field is not present")
	}
	return mti.Pack()
}

func CalculateBitmap(c iso.Component, bm iso.Bitmap) ([]byte, error) {
	if !c.IsComposite() {
		return nil, fmt.Errorf("leaf fields can not generate bitmap")
	}
	return bm.Compute(c)

}

func SortAndPackFields(c iso.Component) ([][]byte, error) {
	if !c.IsComposite() {
		return nil, fmt.Errorf("leaf fields do not have any children")
	}

	fnos := make([]int, len(c.GetChildren()))
	results := make([][]byte, len(c.GetChildren()))

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
