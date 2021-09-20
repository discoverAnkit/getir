package handler

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type RoutesTestSuite struct {
	suite.Suite
	keyValueHandler     *KeyValueHandler
	mongoRequestHandler *MongoRequestHandler
	router              *Router
}

func (s *RoutesTestSuite) SetupTest() {
	//No need to inject repositories in these objects as these tests are meant to only test routing
	//and thus in these tests we wont reach those points where repo will get called
	s.keyValueHandler = &KeyValueHandler{}
	s.mongoRequestHandler = &MongoRequestHandler{}
	s.router = NewRouter(s.keyValueHandler,s.mongoRequestHandler)
}

func Test_RoutesTestSuite(t *testing.T) {
	tests := new(RoutesTestSuite)
	suite.Run(t, tests)
}

func (s *RoutesTestSuite) Test_MethodNotAllowed() {

	req, _ := http.NewRequest(http.MethodPut, "https://api-stage.something.com/getValue", nil)
	res := httptest.NewRecorder()

	s.router.Handle(res,req)

	s.Equal(http.StatusMethodNotAllowed, res.Code)
}

func (s *RoutesTestSuite) Test_RouteNotFound() {

	req, _ := http.NewRequest(http.MethodPost, "https://api-stage.something.com/getValue", nil)
	res := httptest.NewRecorder()

	s.router.Handle(res,req)

	s.Equal(http.StatusNotFound, res.Code)
}

func (s *RoutesTestSuite) Test_ValidGetRoute() {

	req, _ := http.NewRequest(http.MethodGet, "https://api-stage.something.com/getValue", nil)
	res := httptest.NewRecorder()

	s.router.Handle(res,req)

	//API call happens successfully (its a bad request but route is correct, which is the test)
	s.Equal(http.StatusBadRequest, res.Code)
}

func (s *RoutesTestSuite) Test_ValidPostRoute() {

	req, _ := http.NewRequest(http.MethodPost, "https://api-stage.something.com/setKeyValue", strings.NewReader(""))
	res := httptest.NewRecorder()

	s.router.Handle(res,req)

	//API call happens successfully (its a bad request but route is correct, which is the test)
	s.Equal(http.StatusBadRequest, res.Code)
}