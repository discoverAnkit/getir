package handler

import (
	"context"
	"encoding/json"
	"github.com/discoverAnkit/getir/contract"
	"github.com/discoverAnkit/getir/repository"
	"log"
	"net/http"
	"time"
)

const timeFormatLayout = "2006-01-02"

//response codes
const (
	successCode = 0
	badRequestErrorCode = 400
	internalServerErrorCode = 500
)

//response messages
const (
	successMsg = "Success"
	badRequestMsg = "Bad Request"
	internalServerErrorMsg = "Something went wrong!"
)

type MongoRequestHandler struct{
	MongoRepo repository.MongoClient
}

func (h *MongoRequestHandler) GetKeyValueRecords(ctx context.Context,w http.ResponseWriter, r *http.Request) {

	log.Println("Handling GetKeyValueRecords")

	setContentTypeAsJson(w)

	//As per specifications, this api always returns 200 as error codes are embedded in the response object
	response := &contract.GetKeyValueRecordsResponse{}
	responseRecords := make([]contract.Record,0)

	//parse req body
	req := contract.GetKeyValueRecordsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err!=nil {
		log.Println("invalid json, could not be decoded")
		response = prepareResponse(badRequestErrorCode,badRequestMsg,responseRecords)
		resp, _:= json.Marshal(response)
		w.Write(resp)
		return
	}

	//convert input dates into time
	startTime, err := time.Parse(timeFormatLayout, req.StartDate)
	if err!= nil {
		log.Println("start date could not be parsed")
		response = prepareResponse(badRequestErrorCode,badRequestMsg,responseRecords)
		resp, _:= json.Marshal(response)
		w.Write(resp)
		return
	}
	endTime, err := time.Parse(timeFormatLayout, req.EndDate)
	if err!= nil {
		log.Println("end date could not be parsed")
		response = prepareResponse(badRequestErrorCode,badRequestMsg,responseRecords)
		resp, _:= json.Marshal(response)
		w.Write(resp)
		return
	}

	//call repo
	records, err := h.MongoRepo.GetRecordsByCreationTime(ctx,startTime,endTime)
	if err!= nil {
		log.Println("Repository Error calling GetRecordsByCreationTime")
		response = prepareResponse(internalServerErrorCode,internalServerErrorMsg,responseRecords)
		resp, _:= json.Marshal(response)
		w.Write(resp)
		return
	}

	responseRecords = filterRecordsByTotalCount(records,req.MinCount,req.MaxCount)
	response = prepareResponse(successCode,successMsg,responseRecords)
	resp, _:= json.Marshal(response)
	w.Write(resp)
}

func prepareResponse (code int, message string, records []contract.Record) *contract.GetKeyValueRecordsResponse {
	return &contract.GetKeyValueRecordsResponse{
		Code:code,
		Message: message,
		Records: records,
	}
}

//Select only those records for which totalCount (sum of counts in Counts array)
//is in between max count and min count
func filterRecordsByTotalCount(records []contract.KVRecord, minCount, maxCount int) []contract.Record {

	filteredRecords := make([]contract.Record,0)
	for _, record := range records {
		sumCounts := sumArrayElements(record.Counts)
		if sumCounts > minCount && sumCounts < maxCount {
			filteredRecords = append(filteredRecords, contract.Record{
				Key: record.Key,
				TotalCount: sumCounts,
				CreatedAt: record.CreatedAt,
			})
		}
	}
	return filteredRecords
}

//Sum of elements of an integer array
func sumArrayElements (arr []int) int {
	sum := 0
	for i:=0; i<len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}