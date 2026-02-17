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
	fmt.Printf("Test %v\n", b)
}

func TestSuiteExampleAdd(t *testing.T) {
	x := 20
	y := &x
	fmt.Printf("Test %v\n", y)
}
