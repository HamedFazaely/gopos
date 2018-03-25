package iso

//Field is an iso field
type Field struct {
	childField
	baseComponent
	packer Packer
}

//NewISOField is an ISOField factory
func NewISOField(pckr Packer) *Field {
	return &Field{
		packer: pckr,
	}
}

//Pack packs the field
func (f *Field) Pack() ([]byte, error) {
	return f.packer.Pack(f)
}
