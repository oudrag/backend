package cqrs

import "log"

type EventBus struct {
	listeners []Listener
	msgCh     chan *Message
	errCh     chan error
}

func NewEventBus() *EventBus {
	bus := &EventBus{
		listeners: make([]Listener, 0),
		msgCh:     make(chan *Message),
		errCh:     make(chan error),
	}

	go bus.consume(bus.errCh)
	go handleErrors(bus.errCh)

	return bus
}

func (b *EventBus) Publish(msg *Message) error {
	if msg.messageType != QueryMessage {
		return InvalidMessageErr
	}

	b.msgCh <- msg

	return nil
}

func (b *EventBus) Subscribe(listener []Listener) {
	b.listeners = append(b.listeners, listener...)
}

func (b *EventBus) consume(errCh chan<- error) {
	for {
		select {
		case msg := <-b.msgCh:
			for _, listener := range b.listeners {
				if !listener.Interested(msg) {
					continue
				}

				if err := listener.Listen(msg); err != nil {
					errCh <- err
				}
			}
		}
	}
}

func handleErrors(errCh <-chan error) {
	for err := range errCh {
		// TODO: add sentry log here
		log.Println(err)
	}
}
