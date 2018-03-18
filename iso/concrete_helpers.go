package iso

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

type baseComponent struct {
	FieldNo   int
	MaxLength int
	Value     string
	IsComp bool
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
