package cqrs

type MessageType int

const (
	CommandMessage MessageType = iota
	QueryMessage
)

type Message struct {
	Name        string
	payload     Payload
	messageType MessageType
}
