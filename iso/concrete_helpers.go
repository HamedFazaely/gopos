package iso

import (
	"github.com/HamedFazaely/gopos/iso/logger"
)

//childField provides primary functionality and fields of a childField
type childField struct{}

func (cf *childField) AddComponent(fno int, c Component) bool {
	panic("A child filed has no children")
}

func (cf *childField) RemoveComponent(fno int) bool {
	panic("A child filed has no children")
}

func (cf *childField) GetComponent(fno int) (Component, bool) {
	panic("A child filed has no children")
}

func (cf *childField) GetChildren() map[int]Component {
	panic("A child filed has no children")
}

//baseComponent provides primary functionality and fields of an ISOField
type baseComponent struct {
	FieldNo   int
	MaxLength int
	Value     string
	IsComp    bool
	logger.Logger
	packer Packer
}

func (b *baseComponent) GetKey() int {
	return b.FieldNo
}

func (b *baseComponent) GetValue() string {
	return b.Value
}

func (b *baseComponent) GetMaxLength() int {
	return b.MaxLength
}

func (b *baseComponent) IsComposite() bool {
	return b.IsComp
}
