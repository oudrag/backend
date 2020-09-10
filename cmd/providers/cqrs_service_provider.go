package providers

import (
	"github.com/oudrag/server/internal/platform/app"
	"github.com/oudrag/server/internal/platform/cqrs"
	"github.com/oudrag/server/internal/usecase/queries"
)

type CQRSServiceProvider struct{}

func (s CQRSServiceProvider) Boot(container app.Container) error {
	var bus *cqrs.Bus
	if err := container.MakeInto(app.CQRSBusBinding, &bus); err != nil {
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

func (s CQRSServiceProvider) Register(binder app.Binder) {
	binder.Singleton(app.CQRSBusBinding, registerBus)
	binder.Singleton(app.CQRSCommandsBinding, registerCommands)
	binder.Singleton(app.CQRSQueriesBinding, registerQueries)
}

func registerBus(_ app.Container) (interface{}, error) {
	bus := cqrs.NewBus()
	return bus, nil
}

func registerCommands(_ app.Container) (interface{}, error) {
	return map[string]cqrs.Handler{}, nil
}

func registerQueries(_ app.Container) (interface{}, error) {
	return map[string]cqrs.Handler{
		queries.FetchTodayEventsQuery: new(queries.FetchTodayEvents),
	}, nil
}

func bootCommands(c app.Container, bus *cqrs.Bus) error {
	var commandHandlers map[string]cqrs.Handler
	if err := c.MakeInto(app.CQRSCommandsBinding, &commandHandlers); err != nil {
		return err
	}

	for name, handler := range commandHandlers {
		if needInit, ok := handler.(app.HasInit); ok {
			if err := needInit.Init(c); err != nil {
				return err
			}
		}

		bus.AddCommandHandler(name, handler)
	}

	return nil
}

func bootQueries(c app.Container, bus *cqrs.Bus) error {
	var queryHandlers map[string]cqrs.Handler
	if err := c.MakeInto(app.CQRSQueriesBinding, &queryHandlers); err != nil {
		return err
	}

	for name, handler := range queryHandlers {
		if needInit, ok := handler.(app.HasInit); ok {
			if err := needInit.Init(c); err != nil {
				return err
			}
		}

		bus.AddQueryHandler(name, handler)
	}

	return nil
}
