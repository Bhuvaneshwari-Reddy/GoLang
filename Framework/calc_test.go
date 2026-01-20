package framework

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(5, 6)
	expected := 11
	if result != expected {
		t.Errorf("Add (5,6)=%d;expected %d", result, expected)
	}
}

func TestAddWithTestify(t *testing.T) {
	result := Add(2, 4)
	expected := 6

	assert.Equal(t, expected, result, "Add(2, 3) should equal 5")
}

type MyObject struct {
	Value string
}

func TestSuiteExample(t *testing.T) {
	a := 10
	b := &a
	fmt.Printf("Test", b)
}

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
