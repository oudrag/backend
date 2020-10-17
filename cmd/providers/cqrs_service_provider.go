package providers

import (
	"github.com/oudrag/server/internal/core/app"
	"github.com/oudrag/server/internal/core/cqrs"
	"github.com/oudrag/server/internal/usecase/commands"
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

func bootCommands(c app.Container) error {
	var bus *cqrs.CommandBus
	if err := c.MakeInto(app.CQRSCommandBusBinding, &bus); err != nil {
		return err
	}

	bus.RegisterHandlers([]cqrs.Handler{
		new(commands.SignOnUserWithSSOCommandHandler),
	})

	return nil
}

func bootQueries(c app.Container) error {
	var bus *cqrs.QueryBus
	if err := c.MakeInto(app.CQRSQueryBusBinding, &bus); err != nil {
		return err
	}

	bus.RegisterHandlers([]cqrs.Handler{})

	return nil
}

func bootListeners(c app.Container) error {
	var bus *cqrs.EventBus
	if err := c.MakeInto(app.CQRSEventBusBinding, &bus); err != nil {
		return err
	}

	bus.Subscribe([]cqrs.Listener{})

	return nil
}
