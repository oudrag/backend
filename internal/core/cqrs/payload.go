package cqrs

import (
	"context"
	"fmt"
	"reflect"
)

type Payload struct {
	ctx  context.Context
	data hashMap
}

func NewPayload(ctx context.Context) *Payload {
	return &Payload{
		ctx:  ctx,
		data: make(map[string]interface{}),
	}
}

func (p *Payload) Add(key string, value interface{}) *Payload {
	p.data.add(key, value)

	return p
}

func (p *Payload) Get(key string, defaultValue ...interface{}) interface{} {
	return p.data.get(key, defaultValue...)
}

func (p *Payload) GetAs(key string, result interface{}) error {
	value := p.Get(key)
	resultType := reflect.TypeOf(result).Elem()
	resultVal := reflect.ValueOf(result).Elem()
	concreteType := reflect.TypeOf(value)

	if !concreteType.AssignableTo(resultType) {
		return fmt.Errorf(
			"cannot assign concreate value (type %v) of %s to given parameter (type %v)",
			concreteType,
			key,
			resultType,
		)
	}

	resultVal.Set(reflect.ValueOf(value))

	return nil
}
