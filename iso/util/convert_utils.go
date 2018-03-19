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
	for _, d := range val {
		n, err := strconv.Atoi(string(d))
		if err != nil {
			return nil, err
		}
		result = append(result, byte(n))
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
