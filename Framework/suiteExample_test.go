package framework

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	VariableStartAtFive int
}

func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableStartAtFive = 5
}

func (suite *ExampleTestSuite) TestExample() {
	suite.Equal(suite.VariableStartAtFive, 5)
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
