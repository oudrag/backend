package routing

import "github.com/gin-gonic/gin"

const (
	RouterBinding         = "routing.router"
	RoutesListBinding     = "routing.routes"
	MiddlewareListBinding = "routing.middleware"
)

// Method is a custom string type for HTTP verbs.
type Method string

func (m Method) String() string {
	return string(m)
}

// List of available HTTP verbs.
const (
	Post    Method = "POST"
	Get     Method = "GET"
	Head    Method = "HEAD"
	Put     Method = "PUT"
	Patch   Method = "PATCH"
	Delete  Method = "DELETE"
	Options Method = "OPTIONS"
)

type Handler interface {
	Handle(ctx *gin.Context)
}
