package iso

import (
	"github.com/HamedFazaely/gopos/iso/logger"
)

//Message is an iso message
type Message struct {
	baseComponent
	children map[int]Component
}

//NewISOMsg is an ISOMessage factory
func NewISOMsg(pcker Packer, lgr logger.Logger) *Message {
	return &Message{
		children: make(map[int]Component),
		baseComponent: baseComponent{
			packer: pcker,
			IsComp: true,
			Logger: lgr,
		},
	}
}

//AddComponent adds a field to message
func (m *Message) AddComponent(fno int, c Component) bool {
	m.Log(logger.Trace, "adding field no %d to iso message", fno)
	_, exists := m.children[fno]

	if exists { // field exists already, just return false
		return false
	}

	m.children[fno] = c

	return true

}

//RemoveComponent removes field by number
func (m *Message) RemoveComponent(fno int) bool {
	m.Log(logger.Trace, "removing field no : %d", fno)
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
	m.Log(logger.Trace, "packing iso message")
	return m.packer.Pack(m)
}
