package cqrs

type CommandBus struct {
	handlers map[string]Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]Handler),
	}
}

func (b *CommandBus) RegisterHandler(name string, handler Handler) {
	b.handlers[name] = handler
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
