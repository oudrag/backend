package cqrs

import (
	"fmt"
	"reflect"
)

type MessageType int

const (
	CommandMessage MessageType = iota
	QueryMessage
	EventMessage
)

type Message struct {
	Name        string
	payload     interface{}
	messageType MessageType
}

func (m *Message) Payload() interface{} {
	return m.payload
}

func (m *Message) PayloadAs(i interface{}) error {
	pt := reflect.TypeOf(m.payload)
	t := reflect.TypeOf(i).Elem()

	if !pt.AssignableTo(t) {
		return fmt.Errorf("cannot assign payload (%v) to given variable (%v)", pt, t)
	}

	reflect.ValueOf(i).Elem().Set(reflect.ValueOf(m.payload))
	return nil
}

func NewCommand(cmd Command) *Message {
	return &Message{
		Name:        cmd.GetCommandName(),
		payload:     cmd,
		messageType: CommandMessage,
	}
}

func NewQuery(q Query) *Message {
	return &Message{
		Name:        q.GetQueryName(),
		payload:     q,
		messageType: QueryMessage,
	}
}

func NewEvent(e Event) *Message {
	return &Message{
		Name:        e.GetEventName(),
		payload:     e,
		messageType: EventMessage,
	}
}
