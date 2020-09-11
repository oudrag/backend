package routing

import "github.com/gin-gonic/gin"

// Method is a custom string type for HTTP verbs.
type Method string

func (m Method) String() string {
	return string(m)
}

// List of available HTTP verbs.
const (
	headMethod    Method = "HEAD"
	optionsMethod Method = "OPTIONS"
	getMethod     Method = "GET"
	postMethod    Method = "POST"
	putMethod     Method = "PUT"
	patchMethod   Method = "PATCH"
	deleteMethod  Method = "DELETE"
)

type Handler interface {
	Handle(ctx *gin.Context)
}
