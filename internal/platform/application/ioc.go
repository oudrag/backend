package application

import (
	"fmt"
	"reflect"
)

// Resolver is a function that returns a concrete implementation for an
// abstraction key.
type Resolver func(app Container) (interface{}, error)

type binding struct {
	resolver  Resolver
	singleton bool
}

// IoC is the IoC core and config management
type IoC struct {
	Booted    bool
	resolved  map[string]interface{}
	bindings  map[string]binding
	providers []ServiceProvider
}

// Singleton registers a shared binding in the container.
func (a *IoC) Singleton(abstract string, resolver Resolver) {
	a.bindings[abstract] = binding{
		resolver:  resolver,
		singleton: true,
	}
}

// Bind registers a binding in the container.
func (a *IoC) Bind(abstract string, resolver Resolver) {
	a.bindings[abstract] = binding{
		resolver:  resolver,
		singleton: false,
	}
}

// Make returns the concrete object associated with the given abstract key.
// it returns error if no abstraction bounded to ioc or bounded resolver returns
// error.
func (a *IoC) Make(abstract string) (interface{}, error) {
	if resolved, ok := a.resolved[abstract]; ok {
		return resolved, nil
	} else if binding, ok := a.bindings[abstract]; ok {
		concrete, err := binding.resolver(a)
		if err != nil {
			return nil, err
		}

		// if the binding is a singleton binding, save the concrete.
		if binding.singleton {
			a.resolved[abstract] = concrete
		}

		return concrete, nil
	}

	return nil, fmt.Errorf("no binding found for %s", abstract)
}

// MakeInto makes the concrete value associated with the given abstract key and
// assign it to result variable passed by reference. If the concrete value is
// not assignable to the result variable it returns error.
func (a *IoC) MakeInto(abstract string, result interface{}) error {
	concrete, err := a.Make(abstract)
	if err != nil {
		return err
	}

	resultType := reflect.TypeOf(result).Elem()
	resultVal := reflect.ValueOf(result).Elem()
	concreteType := reflect.TypeOf(concrete)

	if !concreteType.AssignableTo(resultType) {
		return fmt.Errorf(
			"cannot assign concreate value (type %v) of %s to given parameter (type %v)",
			concreteType,
			abstract,
			resultType,
		)
	}

	resultVal.Set(reflect.ValueOf(concrete))

	return nil
}

// Boot registers and run the application services.
func (a *IoC) Boot() (err error) {
	if a.Booted {
		return nil
	}

	a.registerServices()
	err = a.runServices()
	a.Booted = true

	return err
}

func (a *IoC) registerServices() {
	for _, provider := range a.providers {
		provider.Register(a)
	}
}

func (a *IoC) runServices() error {
	for _, provider := range a.providers {
		if bootable, isBootable := provider.(BootableServiceProvider); isBootable {
			if err := bootable.Boot(a); err != nil {
				return err
			}
		}
	}

	return nil
}

// NewIoC returns new instance of application.
func NewIoC(providers []ServiceProvider) *IoC {
	app := &IoC{
		bindings:  make(map[string]binding),
		resolved:  make(map[string]interface{}),
		providers: providers,
	}

	return app
}
