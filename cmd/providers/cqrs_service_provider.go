package providers

import (
	"github.com/oudrag/server/internal/platform/application"
	"github.com/oudrag/server/internal/platform/cqrs"
	"github.com/oudrag/server/internal/usecase/queries"
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
	binder.Singleton(cqrs.BusBinding, registerBus)
	binder.Singleton(cqrs.CommandsBinding, registerCommands)
	binder.Singleton(cqrs.QueriesBinding, registerQueries)
}

func registerBus(_ application.Container) (interface{}, error) {
	bus := cqrs.NewBus()
	return bus, nil
}

func registerCommands(_ application.Container) (interface{}, error) {
	return map[string]cqrs.Handler{}, nil
}

func registerQueries(_ application.Container) (interface{}, error) {
	return map[string]cqrs.Handler{
		queries.FetchTodayEventsQuery: new(queries.FetchTodayEvents),
	}, nil
}

func bootCommands(c application.Container, bus *cqrs.Bus) error {
	var commandHandlers map[string]cqrs.Handler
	if err := c.MakeInto(cqrs.CommandsBinding, &commandHandlers); err != nil {
		return err
	}

	for name, handler := range commandHandlers {
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
	var queryHandlers map[string]cqrs.Handler
	if err := c.MakeInto(cqrs.QueriesBinding, &queryHandlers); err != nil {
		return err
	}

	for name, handler := range queryHandlers {
		if needInit, ok := handler.(application.HasInit); ok {
			if err := needInit.Init(c); err != nil {
				return err
			}
		}

		bus.AddQueryHandler(name, handler)
	}

	return nil
}
