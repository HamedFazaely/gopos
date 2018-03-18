package misc

import (
	"fmt"
	"testing"
)

func TestBitmap64_SetBit(t *testing.T) {

	m := make([]byte, 16)

	setBit(10, m)
	setBit(120, m)
	setBit(45, m)
	setBit(8,m)
	setBit(16,m)
	setBit(128,m)


	for _, b := range m {
		fmt.Printf("%08b ", b)
	}

}
