package supervisor

import (
	"encoding/json"
	"fmt"
)

type (
	Spec struct {
		super *Supervisor

		meta       *MetaSpec
		objectSpec interface{}
	}

	MetaSpec struct {
		Name    string `json:"name" jsonschema:"required"`
		Kind    string `json:"kind" jsonschema:"required"`
		Version string `json:"version" jsonschema:"required"`
	}
)

const (
	DefaultSpecVersion = "gbase_gcm_v1"
)

// NewSpec creates a new spec from config in json format.
func (s *Supervisor) NewSpec(config string) (spec *Spec, err error) {
	spec = &Spec{
		super: s,
	}

	buff := []byte(config)

	meta := &MetaSpec{Version: DefaultSpecVersion}
	err = json.Unmarshal(buff, meta)
	if err != nil {
		return nil, err
	}
	spec.meta = meta

	object, ok := objectRegistry[meta.Kind]
	if !ok {
		return nil, fmt.Errorf("can not find object (%s)", meta.Kind)
	}
	objectSpec := object.DefaultSpec()
	err = json.Unmarshal(buff, objectSpec)
	if err != nil {
		return nil, err
	}
	spec.objectSpec = objectSpec

	return spec, nil
}
