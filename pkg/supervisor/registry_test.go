package supervisor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	mockObject struct {
	}
)

func (m *mockObject) Kind() string {
	return "mock_object"
}

func TestRegister(t *testing.T) {
	assert.NotPanics(
		t,
		func() { Register(&mockObject{}) },
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
