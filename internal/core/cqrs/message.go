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

func NewQuery(name string, p *Payload) *Message {
	return &Message{
		Name:        name,
		payload:     p,
		messageType: QueryMessage,
	}
}

func NewCommand(name string, p *Payload) *Message {
	return &Message{
		Name:        name,
		payload:     p,
		messageType: CommandMessage,
	}
}

func NewEvent(name string, p *Payload) *Message {
	return &Message{
		Name:        name,
		payload:     p,
		messageType: EventMessage,
	}
}
