package supervisor

import "fmt"

type (
	ObjectEntity struct {
		super *Supervisor

		instance Object
	}
)

// Instance returns the instance of the Object.
func (e *ObjectEntity) Instance() Object {
	return e.instance
}

// NewObjectEntity creates a new ObjectEntity.
func (s *Supervisor) NewObjectEntity(kind string) (*ObjectEntity, error) {
	if _, ok := objectRegistry[kind]; !ok {
		return nil, fmt.Errorf("object (%s) does not exist", kind)
	}

	instance := objectRegistry[kind].Clone()
	return &ObjectEntity{
		super:    s,
		instance: instance,
	}, nil
}
