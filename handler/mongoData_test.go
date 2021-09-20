package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/discoverAnkit/getir/contract"
	"github.com/discoverAnkit/getir/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type MongoRequestHandlerTestSuite struct {
	suite.Suite
	MongoRepo           *mocks.MongoClient
	mongoRequestHandler *MongoRequestHandler
	router              *Router
}

func (s *MongoRequestHandlerTestSuite) SetupTest() {
	//No need to inject repositories in these objects as these tests are meant to only test routing
	//and thus in these tests we wont reach those points where repo will get called
	s.MongoRepo = new(mocks.MongoClient)
	s.mongoRequestHandler = &MongoRequestHandler{
		MongoRepo: s.MongoRepo,
	}
	s.router = NewRouter(nil,s.mongoRequestHandler)
}

func Test_MongoRequestHandlerTestSuite(t *testing.T) {
	tests := new(MongoRequestHandlerTestSuite)
	suite.Run(t, tests)
}

func (s *MongoRequestHandlerTestSuite) Test_GetKeyValueRecords_BadRequestJson() {
	//Given
	req, _ := http.NewRequest(http.MethodPost, "https://api-stage.something.com/getKeyValueRecords",strings.NewReader(`{\"startDate\":\"2015-01-01\",\"endDate\":\"2015-01-10\",\"minCount\":2500,\"maxCount\":\"3000\"}`))
	res := httptest.NewRecorder()

	//when
	s.router.Handle(res,req)

	//then
	var apiResp contract.GetKeyValueRecordsResponse
	json.Unmarshal(res.Body.Bytes(),&apiResp)
	s.Equal(http.StatusOK, res.Code)
	s.Equal(apiResp.Code,badRequestErrorCode)
	s.Equal(apiResp.Message,badRequestMsg)
	s.Equal(len(apiResp.Records),0)
}

func (s *MongoRequestHandlerTestSuite) Test_GetKeyValueRecords_InvalidDate() {
	//Given
	reqPayload := contract.GetKeyValueRecordsRequest{
		StartDate: "abc",
		EndDate: "def",
		MinCount:12,
		MaxCount: 134,
	}
	reqBody, _ := json.Marshal(reqPayload)
	req, _ := http.NewRequest(http.MethodPost, "https://api-stage.something.com/getKeyValueRecords",bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()

	//when
	s.router.Handle(res,req)

	//then
	var apiResp contract.GetKeyValueRecordsResponse
	json.Unmarshal(res.Body.Bytes(),&apiResp)
	s.Equal(http.StatusOK, res.Code)
	s.Equal(apiResp.Code,badRequestErrorCode)
	s.Equal(apiResp.Message,badRequestMsg)
	s.Equal(len(apiResp.Records),0)
}

func (s *MongoRequestHandlerTestSuite) Test_GetKeyValueRecords_DBError() {
	//Given
	reqPayload := contract.GetKeyValueRecordsRequest{
		StartDate: "2020-01-01",
		EndDate: "2021-01-01",
		MinCount:12,
		MaxCount: 134,
	}
	reqBody, _ := json.Marshal(reqPayload)
	req, _ := http.NewRequest(http.MethodPost, "https://api-stage.something.com/getKeyValueRecords",bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()

	startTime, _ := time.Parse(timeFormatLayout, reqPayload.StartDate)
	endTime, _ := time.Parse(timeFormatLayout, reqPayload.EndDate)
	s.MongoRepo.On("GetRecordsByCreationTime",mock.Anything,startTime,endTime).Return(nil,errors.New("something"))

	//when
	s.router.Handle(res,req)

	//then
	var apiResp contract.GetKeyValueRecordsResponse
	json.Unmarshal(res.Body.Bytes(),&apiResp)
	s.Equal(http.StatusOK, res.Code)
	s.Equal(apiResp.Code,internalServerErrorCode)
	s.Equal(apiResp.Message,internalServerErrorMsg)
	s.Equal(len(apiResp.Records),0)
}

func (s *MongoRequestHandlerTestSuite) Test_GetKeyValueRecords_Success() {
	//Given
	reqPayload := contract.GetKeyValueRecordsRequest{
		StartDate: "2020-01-01",
		EndDate: "2021-01-01",
		MinCount:12,
		MaxCount: 134,
	}
	reqBody, _ := json.Marshal(reqPayload)
	req, _ := http.NewRequest(http.MethodPost, "https://api-stage.something.com/getKeyValueRecords",bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()

	startTime, _ := time.Parse(timeFormatLayout, reqPayload.StartDate)
	endTime, _ := time.Parse(timeFormatLayout, reqPayload.EndDate)
	loc,_ := time.LoadLocation("Asia/Jerusalem")
	createdAt := primitive.NewDateTimeFromTime(time.Date(2020,time.May,5,0,0,0,0,loc)) //5th May, 2020 say for given input dates, all the mongo records were created on this date
	expectedMongoRecords := []contract.KVRecord{
		{
			ID: primitive.ObjectID{},
			Key: "abc",
			Value: "def",
			CreatedAt: createdAt,
			Counts: []int{1,12,13},
		},
		{
			ID: primitive.ObjectID{},
			Key: "abc",
			Value: "cvg",
			CreatedAt: createdAt,
			Counts: []int{1,2,3},
		},
		{
			ID: primitive.ObjectID{},
			Key: "def",
			Value: "cvg",
			CreatedAt: createdAt,
			Counts: []int{1,12,13},
		},
	}
	expectedApiResponse := contract.GetKeyValueRecordsResponse{
		Code: successCode,
		Message: successMsg,
		Records: []contract.Record{
			{
				Key: "abc",
				CreatedAt: createdAt,
				TotalCount: 26,
			},
			{
				Key: "def",
				CreatedAt: createdAt,
				TotalCount: 26,
			},
		},
	}
	s.MongoRepo.On("GetRecordsByCreationTime",mock.Anything,startTime,endTime).Return(expectedMongoRecords,nil)

	//when
	s.router.Handle(res,req)

	//then
	var apiResp contract.GetKeyValueRecordsResponse
	json.Unmarshal(res.Body.Bytes(),&apiResp)
	s.Equal(http.StatusOK, res.Code)
	s.Equal(apiResp,expectedApiResponse)
}