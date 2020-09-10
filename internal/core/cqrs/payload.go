package cqrs

import "context"

type Payload struct {
	ctx  context.Context
	data map[string]interface{}
}

func NewPayload(ctx context.Context) *Payload {
	return &Payload{
		ctx:  ctx,
		data: make(map[string]interface{}),
	}
}

func (p *Payload) Add(key string, value interface{}) *Payload {
	p.data[key] = value

	return p
}

func (p *Payload) GetAs(key string, v interface{}) error {
	panic("implement me")
}
