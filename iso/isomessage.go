package iso



//Message is an iso message
type Message struct {
	baseComponent
	children map[int]Component
	packer   Packer
}

//SetPacker sets a packer for the message
func (m *Message) SetPacker(p Packer) {
	m.packer = p
}

//AddComponent adds a field to message
func (m *Message) AddComponent(fno int, c Component) bool {
	initMessageIfNot(m)
	_, exists := m.children[fno]

	if exists { // field exists already, just return false
		return false
	}

	m.children[fno] = c

	return true

}

//RemoveComponent removes field by number
func (m *Message) RemoveComponent(fno int) bool {
	initMessageIfNot(m)
	_, exists := m.children[fno]
	if exists {
		delete(m.children, fno)
	}
	return exists

}

//GetComponent ...
func (m *Message) GetComponent(fno int) (Component, bool) {
	initMessageIfNot(m)
	c, exists := m.children[fno]
	return c, exists
}

//GetChildren ...
func (m *Message) GetChildren() map[int]Component {
	initMessageIfNot(m)
	return m.children
}

//Pack ...
func (m *Message) Pack() ([]byte, error) {
	initMessageIfNot(m)
	return m.packer.Pack(m)
}


func initMessageIfNot(m *Message) {
	if m.children == nil {
		m.children = make(map[int]Component)
	}
}
