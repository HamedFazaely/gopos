package util

import (
	"strconv"

	"github.com/HamedFazaely/gopos/iso/ierrors"
)

func splitNumberString(val string) ([]byte, error) {
	if !IsASCII(val) {
		return nil, ierrors.ErrNoneASCIIValue
	}
	result := make([]byte, len(val))
	for i := 0; i < len(val); i++ {
		n, err := strconv.Atoi(string(val[i]))
		if err != nil {
			return nil, err
		}
		result[i] = byte(n)
	}
	return result, nil
}

//IsASCII checks a string to see if it only containd ASCII chars
func IsASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

//ASCIIToBCD converts a digit string to its BCD equivalent
//if number of digits are odd the result will contain a 0000 nimble to its right most position
func ASCIIToBCD(num string) ([]byte, error) {
	spl, err := splitNumberString(num)
	if err != nil {
		return nil, err
	}
	var res []byte
	for i := 0; i < len(spl); i += 2 {
		bcd := spl[i] << 4
		if i < len(spl)-1 {
			bcd |= spl[i+1]
		}
		res = append(res, bcd)
	}
	return res, nil
}

//IsNumericString checks a string to see if it only contains digits
func IsNumericString(s string) bool {
	for _, b := range s {
		if b < '0' || b > '9' {
			return false
		}
	}
	return true
}
