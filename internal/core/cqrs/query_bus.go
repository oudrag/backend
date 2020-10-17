package cqrs

import (
	"reflect"
	"strings"
)

type QueryBus struct {
	handlers map[string]Handler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]Handler),
	}
}

func (b *QueryBus) RegisterHandlers(handlers []Handler) {
	for _, handler := range handlers {
		commandName := strings.Replace(
			reflect.TypeOf(handler).Elem().Name(),
			"Handler", "", 1,
		)
		b.handlers[commandName] = handler
	}
}

func (b *QueryBus) Dispatch(msg *Message) Response {
	if msg.messageType != QueryMessage {
		return ErrResponse(InvalidMessageErr)
	}

	for name, handler := range b.handlers {
		if msg.Name == name {
			return handler.Handle(msg)
		}
	}

	return ErrResponse(HandlerNotFoundErr)
}
