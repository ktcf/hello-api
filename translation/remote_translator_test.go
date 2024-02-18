package translation_test

import (
	"errors"
	"github.com/ktcf/hello-api/translation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RemoteServiceTestSuite struct {
	suite.Suite
	client    *MockHelloClient
	underTest *translation.RemoteService
}

type MockHelloClient struct {
	mock.Mock
}

func TestRemoteServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RemoteServiceTestSuite))
}

func (m *MockHelloClient) Translate(word string, language string) (string, error) {
	args := m.Called(word, language)
	return args.String(0), args.Error(1)
}

func (suite *RemoteServiceTestSuite) SetupTest() {
	suite.client = new(MockHelloClient)
	suite.underTest = translation.NewRemoteService(suite.client)
}

func (suite *RemoteServiceTestSuite) TestTranslate() {
	suite.client.On("Translate", "foo", "bar").Return("baz", nil)

	res := suite.underTest.Translate("foo", "bar")

	suite.Equal("baz", res)
	suite.client.AssertExpectations(suite.T())
}

func (suite *RemoteServiceTestSuite) TestTranslate_Error() {
	suite.client.On("Translate", "foo", "bar").Return("", errors.New("failure"))

	res := suite.underTest.Translate("foo", "bar")

	suite.Equal("", res)
	suite.client.AssertExpectations(suite.T())
}

func (suite *RemoteServiceTestSuite) TestTranslate_Cache() {
	suite.client.On("Translate", "foo", "bar").Return("baz", nil).Once()

	res1 := suite.underTest.Translate("foo", "bar")
	res2 := suite.underTest.Translate("Foo", "bar")

	suite.Equal("baz", res1)
	suite.Equal("baz", res2)
	suite.client.AssertExpectations(suite.T())
}
