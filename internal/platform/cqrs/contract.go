package cqrs

const (
	BusBinding      = "cqrs.bus"
	CommandsBinding = "cqrs.commands"
	QueriesBinding  = "cqrs.queries"
)

type Handler interface {
	Handle(cmd *Message) Response
}

// Response determines command dispatching return type
type Response struct {
	err    error
	result interface{}
}

// Ok determines executing command was succeed
func (r Response) Ok() bool {
	return r.err == nil
}

// Err returns error instance when command execution was failed
func (r Response) Err() error {
	return r.err
}

// Result returns command execution returned data
// By default it is nil
func (r Response) Result() interface{} {
	return r.result
}

// HasResult determines the command execution response
// Has any result when succeed
func (r Response) HasResult() bool {
	return r.result != nil
}

// ErrResponse creates new instance of Failed command execution response
func ErrResponse(reason error) Response {
	return Response{err: reason}
}

// OkResponse creates new instance of succeed command execution response
func OkResponse(result interface{}) Response {
	return Response{result: result}
}
