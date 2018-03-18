package iso

//Field is an iso field
type Field struct {
	childField
	baseComponent
	packer Packer
}

//SetPacker sets a packer for this field
func (f *Field) SetPacker(p Packer) {
	f.packer = p
}

//Pack packs the field
func (f *Field) Pack() ([]byte, error) {
	return f.packer.Pack(f)
}
