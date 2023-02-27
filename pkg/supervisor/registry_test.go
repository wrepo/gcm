package supervisor

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	mockObject struct {
	}
)

var (
	once sync.Once
)

func (m *mockObject) Kind() string {
	return "mock_object"
}

// Clone returns a copy of the Object.
func (m *mockObject) Clone() Object {
	return &mockObject{}
}

func TestRegister(t *testing.T) {
	assert.NotPanics(
		t,
		func() { once.Do(func() { Register(&mockObject{}) }) },
		"panic in Register",
	)
	assert.Equal(t, 1, len(objectRegistry))
	_, ok := objectRegistry["mock_object"]
	assert.True(t, ok)
	assert.PanicsWithValue(
		t,
		"object (mock_object) already exists",
		func() { Register(&mockObject{}) },
	)
}

func TestSupervisor_NewObjectEntity(t *testing.T) {
	once.Do(func() { Register(&mockObject{}) })
	sv := NewSupervisor(DefaultConfig())
	o, err := sv.NewObjectEntity("mock_object")
	assert.Nil(t, err)
	assert.Equal(t, &mockObject{}, o.Instance())
}
