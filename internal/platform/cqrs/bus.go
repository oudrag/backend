package cqrs

import (
	"fmt"
)

var (
	HandlerNotFoundErr = fmt.Errorf("no handler found for this message")
	InvalidMessageErr  = fmt.Errorf("invalid message passed to bus")
)

type Bus struct {
	commands map[string]Handler
	queries  map[string]Handler
}

func NewBus() *Bus {
	return &Bus{
		commands: make(map[string]Handler),
		queries:  make(map[string]Handler),
	}
}

func (b *Bus) AddCommandHandler(name string, handler Handler) {
	b.commands[name] = handler
}

func (b *Bus) AddQueryHandler(name string, handler Handler) {
	b.queries[name] = handler
}

func (b *Bus) Dispatch(msg *Message) Response {
	switch msg.messageType {
	case CommandMessage:
		return b.runCommandHandler(msg)
	case QueryMessage:
		return b.runQueryHandler(msg)
	default:
		return ErrResponse(InvalidMessageErr)
	}
}

func (b *Bus) runCommandHandler(msg *Message) Response {
	for name, handler := range b.commands {
		if msg.Name == name {
			return handler.Handle(msg)
		}
	}

	return ErrResponse(HandlerNotFoundErr)
}

func (b *Bus) runQueryHandler(msg *Message) Response {
	for name, handler := range b.queries {
		if msg.Name == name {
			return handler.Handle(msg)
		}
	}

	return ErrResponse(HandlerNotFoundErr)
}
