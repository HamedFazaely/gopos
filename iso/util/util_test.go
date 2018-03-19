package util

import (

	"testing"
)

func TestASCIIToBCD(t *testing.T) {

	var tests = []struct {
		input string
		want  []byte
	}{
		{"1234", []byte{0x12, 0x34}},
		{"7675701", []byte{0x76, 0x75, 0x70, 0x10}},
	}

	for _, test := range tests {
		got, err := ASCIIToBCD(test.input)
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < len(got); i++ {
			if got[i] != test.want[i] {
				t.Errorf("got : %08b expected : %08b", got[i], test.want[i])
			}
		}
	}

}
