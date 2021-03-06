package handler

import (
	"context"
	"encoding/json"
	"github.com/discoverAnkit/getir/contract"
	"github.com/discoverAnkit/getir/repository"
	"log"
	"net/http"
)

//error msgs
const (
	missingQueryParamMsg = "Missing Query Param"
	notFoundMsg = "Not Found"
)

type KeyValueHandler struct{
	InMemoryRepository repository.InMemoryClient
}

func (h *KeyValueHandler) SetKeyValue(ctx context.Context,w http.ResponseWriter, r *http.Request) {

	log.Println("Handling SetKeyValue")

	//parse req body
	setKeyValueRequest := contract.SetKeyValueRequest{}
	err := json.NewDecoder(r.Body).Decode(&setKeyValueRequest)
	if err!=nil {
		http.Error(w,badRequestMsg,http.StatusBadRequest)
		return
	}
	log.Printf("Key : %s, Value %s",setKeyValueRequest.Key,setKeyValueRequest.Value)

	//call repo
	err = h.InMemoryRepository.SetKeyValue(ctx,setKeyValueRequest.Key,setKeyValueRequest.Value)
	if err!= nil {
		log.Println("Repository Error calling SetKeyValue")
		http.Error(w,internalServerErrorMsg,http.StatusInternalServerError)
		return
	}

	setKeyValueResponse := contract.SetKeyValueResponse{
		Key: setKeyValueRequest.Key,
		Value: setKeyValueRequest.Value,
	}

	resp, err := json.Marshal(setKeyValueResponse)
	if err!= nil {
		log.Println("Error in JSON Marshaling")
		http.Error(w,internalServerErrorMsg,http.StatusInternalServerError)
		return
	}

	setContentTypeAsJson(w)
	w.Write(resp)
}

func (h *KeyValueHandler) GetValue(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	log.Println("Handling GetValue")

	values :=  r.URL.Query()
	keyValues, found := values["key"]
	if !found {
		log.Println("Request could not be completed as key was missing in query params")
		http.Error(w,missingQueryParamMsg,http.StatusBadRequest)
		return
	}

	value := h.InMemoryRepository.GetValue(ctx,keyValues[0])
	if len(value) == 0 {
		log.Println("Key not found")
		http.Error(w,notFoundMsg,http.StatusNotFound)
		return
	}

	getValueResponse := contract.GetValueResponse{
		Key: keyValues[0],
		Value: value,
	}

	resp, err := json.Marshal(getValueResponse)
	if err!= nil {
		log.Println("Error in JSON Marshaling")
		http.Error(w,internalServerErrorMsg,http.StatusInternalServerError)
		return
	}

	setContentTypeAsJson(w)
	w.Write(resp)
}

func setContentTypeAsJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}