package iso

//Message is an iso message
type Message struct {
	baseComponent
	children map[int]Component
	packer   Packer
}

//NewISOMsg is an ISOMessage factory
func NewISOMsg(pcker Packer) *Message {
	return &Message{
		children: make(map[int]Component),
		packer:   pcker,
	}
}

//AddComponent adds a field to message
func (m *Message) AddComponent(fno int, c Component) bool {
	_, exists := m.children[fno]

	if exists { // field exists already, just return false
		return false
	}

	m.children[fno] = c

	return true

}

//RemoveComponent removes field by number
func (m *Message) RemoveComponent(fno int) bool {
	_, exists := m.children[fno]
	if exists {
		delete(m.children, fno)
	}
	return exists

}

//GetComponent ...
func (m *Message) GetComponent(fno int) (Component, bool) {

	c, exists := m.children[fno]
	return c, exists
}

//GetChildren ...
func (m *Message) GetChildren() map[int]Component {
	return m.children
}

//Pack ...
func (m *Message) Pack() ([]byte, error) {

	return m.packer.Pack(m)
}
