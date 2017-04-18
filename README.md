# go-type-registry

Allows for the registration and use of structs dynamically


Example:

```
func generateRegistry(r *runtime.Registry) error {
	r.Put(reflect.TypeOf(&modules.Gitlab{}))
	return nil
}

registry, err := runtime.NewRegistry(generateRegistry)
if err != nil {
	fmt.Printf("error: %s\n", err.Error())
	return
}

_ = registry.Get("*modules.Gitlab")
  ```
