package cqrs

import (
	"fmt"

	"github.com/oudrag/server/internal/platform/application"
)

var (
	HandlerNotFoundErr = fmt.Errorf("no handler found for this message")
	InvalidMessageErr  = fmt.Errorf("invalid message passed to bus")
)

type Bus struct {
	app      application.Container
	commands map[string]Handler
	queries  map[string]Handler
}

func NewBus(app application.Container) *Bus {
	return &Bus{
		app:      app,
		commands: make(map[string]Handler),
		queries:  make(map[string]Handler),
	}
}

func (b *Bus) RegisterCommands(commands map[string]Handler) {
	b.commands = commands
}

func (b *Bus) RegisterQueries(queries map[string]Handler) {
	b.queries = queries
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
			if needInit, ok := handler.(application.HasInit); ok {
				if err := needInit.Init(b.app); err != nil {
					return ErrResponse(err)
				}
			}

			return handler.Handle(msg)
		}
	}

	return ErrResponse(HandlerNotFoundErr)
}

func (b *Bus) runQueryHandler(msg *Message) Response {
	for name, handler := range b.queries {
		if msg.Name == name {
			if needInit, ok := handler.(application.HasInit); ok {
				if err := needInit.Init(b.app); err != nil {
					return ErrResponse(err)
				}
			}

			return handler.Handle(msg)
		}
	}

	return ErrResponse(HandlerNotFoundErr)
}
