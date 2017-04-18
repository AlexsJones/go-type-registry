package main

import (
	"fmt"

	reg "github.com/AlexsJones/go-type-registry/core"
)

type foo struct {
}

func (*foo) Hello() string {
	return "Hello"
}

func generateRegistry(r *reg.Registry) error {
	r.Put(&foo{})
	return nil
}
func main() {
	registry, err := reg.NewRegistry(generateRegistry)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	value, err := registry.Get("*main.foo")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	//Hitting language limitations here
	fmt.Println(value.Unwrap().(*foo).Hello())
}
