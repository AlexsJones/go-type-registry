package registry

import (
	"errors"
	"log"
	"reflect"
)

//Entry is a utility structure for holding type information for reflection
type Entry struct {
	Name    string
	RefType reflect.Type
}

//Unwrap brings the entry back into a real object
func (e *Entry) Unwrap() interface{} {
	return reflect.New(e.RefType).Elem().Interface()
}

//Registry is a utility structure for holding type information for reflection
type Registry struct {
	storedEntries map[string]*Entry
}

//Put will register a type against a string
func (r *Registry) Put(i interface{}) {
	t := reflect.TypeOf(i)
	e := &Entry{RefType: t, Name: t.String()}
	log.Println("Registering as " + t.String())
	log.Println(reflect.TypeOf(i))
	r.storedEntries[t.String()] = e
}

//Get will return a new instance of type if avialable
func (r *Registry) Get(name string) (*Entry, error) {
	e := r.storedEntries[name]
	if e == nil {
		return nil, errors.New("Not found")
	}
	return e, nil
}

//NewRegistry takes optional functions that can manipulate the object initialised
func NewRegistry(options ...func(*Registry) error) (*Registry, error) {
	r := &Registry{storedEntries: make(map[string]*Entry)}
	for _, op := range options {
		err := op(r)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}
