package iso



//Component represents an iso component
type Component interface {
	AddComponent(fno int, c Component) bool
	RemoveComponent(fno int) bool
	GetComponent(fno int) (Component, bool)
	GetChildren() map[int]Component
	SetPacker(p Packer)
	Pack() ([]byte, error)
	GetKey() int
	GetMaxLength() int
	GetValue() string
	IsComposite() bool
}

//Packer represents an iso packer
type Packer interface {
	Pack(c Component) ([]byte, error)
}


type Bitmap interface {
	Compute(c Component) ([]byte,error)
}