package framework

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MyMockObject struct {
	mock.Mock
}
type MyInterface interface {
	DoSomething(number int) (bool, error)
}

func (m *MyMockObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

func targetFuncThatDoesSomethingWithObj(obj MyInterface) {
	obj.DoSomething(123)
}
func TestSomething1(t *testing.T) {
	testObj := new(MyMockObject)
	testObj.On("DoSomething", 123).Return(true, nil)
	targetFuncThatDoesSomethingWithObj(testObj)
	testObj.AssertExpectations(t)
}
