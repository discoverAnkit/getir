package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/discoverAnkit/getir/contract"
	"github.com/discoverAnkit/getir/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type KeyValueHandlerTestSuite struct {
	suite.Suite
	InMemoryRepository  *mocks.InMemoryClient
	keyValueHandler     *KeyValueHandler
	router              *Router
}

func (s *KeyValueHandlerTestSuite) SetupTest() {
	//No need to inject repositories in these objects as these tests are meant to only test routing
	//and thus in these tests we wont reach those points where repo will get called
	s.InMemoryRepository = new(mocks.InMemoryClient)
	s.keyValueHandler = &KeyValueHandler{
		InMemoryRepository: s.InMemoryRepository,
	}
	s.router = NewRouter(s.keyValueHandler,nil)
}

func Test_KeyValueHandlerTestSuite(t *testing.T) {
	tests := new(KeyValueHandlerTestSuite)
	suite.Run(t, tests)
}

func (s *KeyValueHandlerTestSuite) Test_SetKeyValue_BadRequest() {

	req, _ := http.NewRequest(http.MethodPost, "https://api-stage.something.com/setKeyValue",strings.NewReader(`{\"key\":\"sonammm\",\"value\":1}`))
	res := httptest.NewRecorder()

	s.router.Handle(res,req)

	s.Equal(http.StatusBadRequest, res.Code)
	s.Equal(res.Body.String(),badRequestMsg+"\n")
}

func (s *KeyValueHandlerTestSuite) Test_SetKeyValue_InternalError() {

	//Given
	reqPayload := contract.SetKeyValueRequest{
		Key: "abc",
		Value: "def",
	}
	reqBody, _ := json.Marshal(reqPayload)
	req, _ := http.NewRequest(http.MethodPost, "https://api-stage.something.com/setKeyValue", bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()

	s.InMemoryRepository.On("SetKeyValue",mock.Anything,reqPayload.Key,reqPayload.Value).Return(errors.New("something"))

	//when
	s.router.Handle(res,req)

	//then
	s.Equal(http.StatusInternalServerError, res.Code)
	s.Equal(res.Body.String(),internalServerErrorMsg+"\n")

}

func (s *KeyValueHandlerTestSuite) Test_SetKeyValue_Successful() {

	//Given
	reqPayload := contract.SetKeyValueRequest{
		Key: "abc",
		Value: "def",
	}
	reqBody, _ := json.Marshal(reqPayload)
	req, _ := http.NewRequest(http.MethodPost, "https://api-stage.something.com/setKeyValue", bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()

	s.InMemoryRepository.On("SetKeyValue",mock.Anything,reqPayload.Key,reqPayload.Value).Return(nil)

	//when
	s.router.Handle(res,req)

	//then
	var apiResp contract.GetValueResponse
	json.Unmarshal(res.Body.Bytes(),&apiResp)
	s.Equal(http.StatusOK, res.Code)
	s.Equal(apiResp.Key,reqPayload.Key)
	s.Equal(apiResp.Value,reqPayload.Value)
}

func (s *KeyValueHandlerTestSuite) Test_GetValue_MissingQueryParam() {

	req, _ := http.NewRequest(http.MethodGet, "https://api-stage.something.com/getValue",nil)
	res := httptest.NewRecorder()

	s.router.Handle(res,req)

	s.Equal(http.StatusBadRequest, res.Code)
	s.Equal(res.Body.String(),missingQueryParamMsg+"\n")
}

func (s *KeyValueHandlerTestSuite) Test_GetValue_KeyNotFound() {

	//Given
	keyInput := "abc"
	req, _ := http.NewRequest(http.MethodGet, "https://api-stage.something.com/getValue?key="+keyInput, nil)
	res := httptest.NewRecorder()

	s.InMemoryRepository.On("GetValue",mock.Anything,keyInput).Return("")

	//when
	s.router.Handle(res,req)

	//then
	s.Equal(http.StatusNotFound, res.Code)
}

func (s *KeyValueHandlerTestSuite) Test_GetValue_Successful() {

	//Given
	keyInput := "abc"
	valueOutput := "def"
	req, _ := http.NewRequest(http.MethodGet, "https://api-stage.something.com/getValue?key="+keyInput, nil)
	res := httptest.NewRecorder()

	s.InMemoryRepository.On("GetValue",mock.Anything,keyInput).Return(valueOutput)

	//when
	s.router.Handle(res,req)

	//then
	var apiResp contract.GetValueResponse
	json.Unmarshal(res.Body.Bytes(),&apiResp)
	s.Equal(http.StatusOK, res.Code)
	s.Equal(apiResp.Key,keyInput)
	s.Equal(apiResp.Value,valueOutput)
}