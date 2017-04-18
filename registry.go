package registry

import (
	"log"
	"reflect"
)

//Registry is a utility structure for holding type information for reflection
type Registry struct {
	storedTypes map[string]reflect.Type
}

//Put will register a type against a string
func (r *Registry) Put(t reflect.Type) {
	log.Println("Registering as " + t.String())
	r.storedTypes[t.String()] = t
}

//Get will return a new instance of type if avialable
func (r *Registry) Get(name string) (reflect.Value, reflect.Type) {
	n := reflect.New(r.storedTypes[name].Elem())
	return n, n.Type()
}

//NewRegistry takes optional functions that can manipulate the object initialised
func NewRegistry(options ...func(*Registry) error) (*Registry, error) {
	r := &Registry{storedTypes: make(map[string]reflect.Type)}
	for _, op := range options {
		err := op(r)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}
