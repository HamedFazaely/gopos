package iso

import (
	"github.com/HamedFazaely/gopos/iso/logger"
)

//Field is an iso field
type Field struct {
	childField
	baseComponent
}

//NewISOField is an ISOField factory
func NewISOField(pckr Packer, lgr logger.Logger, fno int, val string) *Field {
	return &Field{
		baseComponent: baseComponent{
			packer:  pckr,
			Logger:  lgr,
			FieldNo: fno,
			Value:   val,
			IsComp:  false,
		},
	}
}

//Pack packs the field
func (f *Field) Pack() ([]byte, error) {
	f.Log(logger.Trace, "packing field no : %d with value : %s", f.FieldNo, f.Value)
	return f.packer.Pack(f)
}
