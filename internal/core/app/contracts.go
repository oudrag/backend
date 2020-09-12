package app

const (
	RedisClientBinding    = "redis.client"
	DBConnectionBinding   = "db.connection"
	GraphServerBinding    = "graph.server"
	CQRSCommandBusBinding = "cqrs.commandBus"
	CQRSQueryBusBinding   = "cqrs.queryBus"
	CQRSEventBusBinding   = "cqrs.eventBus"
	CQRSCommandsBinding   = "cqrs.commands"
	CQRSQueriesBinding    = "cqrs.queries"
	CQRSListenersBinding  = "cqrs.listeners"
	GoogleOAuthBinding    = "google.oauth"
	RouterBinding         = "routing.router"
	RoutesListBinding     = "routing.routes"
	MiddlewareListBinding = "routing.middleware"
)

const (
	AppURL = "APP_URL"

	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBDatabase = "DB_DATABASE"
	DBUsername = "DB_USERNAME"
	DBPassword = "DB_PASSWORD"

	RedisHost     = "REDIS_HOST"
	RedisUsername = "REDIS_USERNAME"
	RedisPassword = "REDIS_PASSWORD"

	GoogleClientID     = "GOOGLE_CLIENT_ID"
	GoogleClientSecret = "GOOGLE_CLIENT_SECRET"
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
