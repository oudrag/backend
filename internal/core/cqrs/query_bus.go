package cqrs

type QueryBus struct {
	handlers map[string]Handler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]Handler),
	}
}

func (b *QueryBus) AddHandler(name string, handler Handler) {
	b.handlers[name] = handler
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
