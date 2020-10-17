package cqrs

import (
	"context"
)

type Payload struct {
	ctx  context.Context
	data interface{}
}

func NewPayload(ctx context.Context) *Payload {
	return &Payload{
		ctx:  ctx,
		data: make(map[string]interface{}),
	}
}
