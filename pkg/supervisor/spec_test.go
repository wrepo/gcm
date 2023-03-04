package supervisor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	mockObjectSpec struct {
		MetaSpec
		Value string `json:"value" jsonschema:"required"`
	}
)

func (m *mockObject) DefaultSpec() interface{} {
	return &mockObjectSpec{MetaSpec{Version: DefaultSpecVersion}, ""}
}

func TestSupervisor_NewSpec_MockObject(t *testing.T) {
	assert.NotPanics(
		t,
		func() { once.Do(func() { Register(&mockObject{}) }) },
		"panic in Register",
	)
	conf := `
			{
				"name": "mock_demo",
				"kind": "mock_object",
				"value": "mock1"
			}
		`

	s := NewSupervisor(DefaultConfig())
	spec, err := s.NewSpec(conf)
	assert.Nil(t, err)
	assert.Equal(t, "mock_demo", spec.meta.Name)
	assert.Equal(t, "mock_object", spec.meta.Kind)
	assert.Equal(t, DefaultSpecVersion, spec.meta.Version)
	assert.NotNil(t, spec.objectSpec)
	mockSpec := spec.objectSpec.(*mockObjectSpec)
	assert.Equal(t, "mock_demo", mockSpec.Name)
	assert.Equal(t, "mock_object", mockSpec.Kind)
	assert.Equal(t, DefaultSpecVersion, mockSpec.Version)
	assert.Equal(t, "mock1", mockSpec.Value)
}
