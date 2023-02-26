package supervisor

import "fmt"

type (
	Object interface {
		Kind() string
	}
)

var (
	objectRegistry = map[string]Object{}
)

func Register(o Object) {
	_, ok := objectRegistry[o.Kind()]
	if ok {
		panic(fmt.Sprintf("object (%s) already exists", o.Kind()))
	}
	objectRegistry[o.Kind()] = o
}
