package translation_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ktcf/hello-api/translation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloClientSuite(t *testing.T) {
	suite.Run(t, new(HelloClientSuite))
}

type HelloClientSuite struct {
	suite.Suite
	mockServerService *MockService
	server            *httptest.Server
	underTest         translation.HelloClient
}

type MockService struct {
	mock.Mock
}

func (m *MockService) Translate(word, language string) (string, error) {
	args := m.Called(word, language)
	return args.String(0), args.Error(1)
}

func (suite *HelloClientSuite) SetupSuite() {
	suite.mockServerService = new(MockService)

	handler := func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		defer func(Body io.ReadCloser) { _ = Body.Close() }(r.Body)

		var m map[string]any
		_ = json.Unmarshal(b, &m)

		word := m["word"].(string)
		language := m["language"].(string)

		resp, err := suite.mockServerService.Translate(word, language)
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
		}

		fmt.Println("Response: ", resp)

		if resp == "" {
			http.Error(w, "error", http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, resp)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	suite.server = httptest.NewServer(mux)
	suite.underTest = translation.NewHelloClient(suite.server.URL)
}

func (suite *HelloClientSuite) TearDownSuite() {
	suite.server.Close()
}

func (suite *HelloClientSuite) SetupTest() {
	suite.mockServerService = new(MockService)
}

func (suite *HelloClientSuite) TestCall() {
	suite.mockServerService.
		On("Translate", "foo", "bar").
		Return(`{"translation":"baz"}`, nil)

	res, err := suite.underTest.Translate("foo", "bar")

	suite.NoError(err)
	suite.Equal("baz", res)
}

func (suite *HelloClientSuite) TestCall_APIError() {
	suite.mockServerService.
		On("Translate", "foo", "bar").
		Return("", errors.New("this is a test"))

	resp, err := suite.underTest.Translate("foo", "bar")

	suite.Equal(err, errors.New("error in api"))
	suite.Equal(resp, "")
}

func (suite *HelloClientSuite) TestCall_InvalidJSON() {
	suite.mockServerService.
		On("Translate", "foo", "bar").
		Return("invalid json", nil)

	resp, err := suite.underTest.Translate("foo", "bar")

	suite.Equal(err, errors.New("unable to decode msg"))
	suite.Equal(resp, "")
}
