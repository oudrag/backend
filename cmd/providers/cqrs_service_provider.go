package providers

import (
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/core/cqrs"
	"github.com/oudrag/server/internal/usecase/queries"
)

type CQRSServiceProvider struct{}

func (s CQRSServiceProvider) Boot(container app.Container) error {
	if err := bootCommands(container); err != nil {
		return err
	}
	if err := bootQueries(container); err != nil {
		return err
	}
	return bootListeners(container)
}

func (s CQRSServiceProvider) Register(binder app.Binder) {
	binder.Singleton(app.CQRSCommandBusBinding, registerCommandBus)
	binder.Singleton(app.CQRSQueryBusBinding, registerQueryBus)
	binder.Singleton(app.CQRSEventBusBinding, registerEventBus)
	binder.Singleton(app.CQRSCommandsBinding, registerCommands)
	binder.Singleton(app.CQRSQueriesBinding, registerQueries)
	binder.Singleton(app.CQRSListenersBinding, registerListeners)
}

func registerCommandBus(_ app.Container) (interface{}, error) {
	bus := cqrs.NewCommandBus()
	return bus, nil
}

func registerQueryBus(_ app.Container) (interface{}, error) {
	bus := cqrs.NewQueryBus()
	return bus, nil
}

func registerEventBus(_ app.Container) (interface{}, error) {
	bus := cqrs.NewEventBus()
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

func registerListeners(_ app.Container) (interface{}, error) {
	return []cqrs.Listener{}, nil
}

func bootCommands(c app.Container) error {
	var bus *cqrs.CommandBus
	if err := c.MakeInto(app.CQRSCommandBusBinding, &bus); err != nil {
		return err
	}

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

		bus.AddHandler(name, handler)
	}

	return nil
}

func bootQueries(c app.Container) error {
	var bus *cqrs.QueryBus
	if err := c.MakeInto(app.CQRSQueryBusBinding, &bus); err != nil {
		return err
	}

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

		bus.AddHandler(name, handler)
	}

	return nil
}

func bootListeners(c app.Container) error {
	var bus *cqrs.EventBus
	if err := c.MakeInto(app.CQRSEventBusBinding, &bus); err != nil {
		return err
	}

	var listeners []cqrs.Listener
	if err := c.MakeInto(app.CQRSListenersBinding, &listeners); err != nil {
		return err
	}

	for _, listener := range listeners {
		if needInit, ok := listener.(app.HasInit); ok {
			if err := needInit.Init(c); err != nil {
				return err
			}
		}

		bus.Subscribe(listener)
	}

	return nil
}
