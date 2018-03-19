package packer

import (
	"github.com/HamedFazaely/gopos/iso"

	"fmt"
)

//FixedLengthASCII packs a fixed length field into ASCII
type FixedLengthASCII struct {

}

//Pack implements Packer interface
func (fla FixedLengthASCII) Pack(c iso.Component) ([]byte, error) {


	if len(c.GetValue()) !=c.GetMaxLength(){
		return nil,fmt.Errorf("length missmatch")
	}

	return []byte(c.GetValue()),nil

}
