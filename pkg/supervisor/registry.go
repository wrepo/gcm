package supervisor

import "fmt"

type (
	Object interface {
		// Kind returns the kind of the object.
		Kind() string

		// Clone returns a copy of the object.
		Clone() Object
	}
)

var (
	// objectRegistry is the registry of objects.
	objectRegistry = map[string]Object{}
)

// Register registers an object.
func Register(o Object) {
	_, ok := objectRegistry[o.Kind()]
	if ok {
		panic(fmt.Sprintf("object (%s) already exists", o.Kind()))
	}
	objectRegistry[o.Kind()] = o
}
