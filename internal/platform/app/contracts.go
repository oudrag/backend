package app

const (
	DBConnectionBinding   = "db.connection"
	GraphServerBinding    = "graph.server"
	CQRSBusBinding        = "cqrs.bus"
	CQRSCommandsBinding   = "cqrs.commands"
	CQRSQueriesBinding    = "cqrs.queries"
	RouterBinding         = "routing.router"
	RoutesListBinding     = "routing.routes"
	MiddlewareListBinding = "routing.middleware"
)

// Container is an interface determines application containers behavior.
type Container interface {
	Make(abstract string) (interface{}, error)
	MakeInto(abstract string, result interface{}) error
}

// Binder is an interface determines application Binder behavior.
type Binder interface {
	Bind(abstract string, resolver Resolver)
	Singleton(abstract string, resolver Resolver)
}

// HasInit interface to construct structs with ioc.Container
type HasInit interface {
	Init(container Container) error
}

// ServiceProvider is an interface determines service providers behavior.
type ServiceProvider interface {
	Register(binder Binder)
}

// BootableServiceProvider makes a service provider bootable
type BootableServiceProvider interface {
	Boot(container Container) error
}
