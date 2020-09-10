package providers

import (
	"github.com/oudrag/server/internal/platform/application"
	"github.com/oudrag/server/internal/platform/cqrs"
)

var (
	commands = map[string]cqrs.Handler{}
	queries  = map[string]cqrs.Handler{}
)

type CQRSServiceProvider struct{}

func (s CQRSServiceProvider) Boot(container application.Container) error {
	var bus *cqrs.Bus
	if err := container.MakeInto(cqrs.BusBinding, &bus); err != nil {
		return err
	}

	if err := bootCommands(container, bus); err != nil {
		return err
	}

	if err := bootQueries(container, bus); err != nil {
		return err
	}

	return nil
}

func (s CQRSServiceProvider) Register(binder application.Binder) {
	binder.Bind(cqrs.BusBinding, func(app application.Container) (interface{}, error) {
		bus := cqrs.NewBus()
		return bus, nil
	})
}

func bootCommands(c application.Container, bus *cqrs.Bus) error {
	for name, handler := range commands {
		if needInit, ok := handler.(application.HasInit); ok {
			if err := needInit.Init(c); err != nil {
				return err
			}
		}

		bus.AddCommandHandler(name, handler)
	}

	return nil
}

func bootQueries(c application.Container, bus *cqrs.Bus) error {
	for name, handler := range queries {
		if needInit, ok := handler.(application.HasInit); ok {
			if err := needInit.Init(c); err != nil {
				return err
			}
		}

		bus.AddQueryHandler(name, handler)
	}

	return nil
}
