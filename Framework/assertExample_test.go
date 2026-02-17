package framework

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	var object *MyObject
	assert.Equal(t, 132, 132, "They should be equal")
	assert.NotEqual(t, 134, 143, "They are not equal")
	assert.Nil(t, object)
	object = &MyObject{Value: "Something"}
	if assert.NotNil(t, object) {
		assert.Equal(t, "Something", object.Value)
	}
}
