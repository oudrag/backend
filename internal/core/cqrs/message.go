package cqrs

type MessageType int

const (
	CommandMessage MessageType = iota
	QueryMessage
	EventMessage
)

type Message struct {
	Name        string
	payload     *Payload
	messageType MessageType
}

func (m *Message) Payload() *Payload {
	return m.payload
}

func (m *Message) GetAs(key string, v interface{}) error {
	return m.payload.GetAs(key, v)
}

func NewQuery(name string, p *Payload) *Message {
	return &Message{
		Name:        name,
		payload:     p,
		messageType: QueryMessage,
	}
}
