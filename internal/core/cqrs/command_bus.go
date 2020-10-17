package cqrs

import (
	"reflect"
	"strings"
)

type CommandBus struct {
	handlers map[string]Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]Handler),
	}
}

func (b *CommandBus) RegisterHandlers(handlers []Handler) {
	for _, handler := range handlers {
		commandName := strings.Replace(
			reflect.TypeOf(handler).Elem().Name(),
			"Handler", "", 1,
		)
		b.handlers[commandName] = handler
	}
}

func (b *CommandBus) Dispatch(msg *Message) Response {
	if msg.messageType != CommandMessage {
		return ErrResponse(InvalidMessageErr)
	}

	for name, handler := range b.handlers {
		if msg.Name == name {
			return handler.Handle(msg)
		}
	}

	return ErrResponse(HandlerNotFoundErr)
}
