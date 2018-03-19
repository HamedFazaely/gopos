package util

import (
	"fmt"
	"strconv"
)

func splitNumberString(val string) ([]byte, error) {
	if !isASCII(val) {
		return nil, fmt.Errorf("%s should only contain ascii characters", val)
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

func isASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

func ASCIIToBCD(num string) ([]byte, error) {
	spl, err := splitNumberString(num)
	if err != nil {
		return nil, err
	}
	var res []byte
	for i := 0; i < len(spl); i+=2 {
		bcd := spl[i] << 4
		if i < len(spl)-1 {
			bcd ^= spl[i+1]
		}
		res = append(res, bcd)
	}
	return res, nil
}
