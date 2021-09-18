package handler

import (
	"context"
	"encoding/json"
	"github.com/discoverAnkit/getir/contract"
	"github.com/discoverAnkit/getir/repository"
	"net/http"
)


type KeyValueHandler struct{
	InMemoryRepository repository.InMemoryClient
}

func (h *KeyValueHandler) SetKeyValue(ctx context.Context,w http.ResponseWriter, r *http.Request) {
	//TODO below
	//add logs
	//input validation ? - not really required
	//parse req body
	//call repo
	//return response
	//http error codes
	//Unit Tests

	setKeyValueRequest := contract.SetKeyValueRequest{}
	json.NewDecoder(r.Body).Decode(&setKeyValueRequest)

	//handler error
	h.InMemoryRepository.SetKeyValue(ctx,setKeyValueRequest.Key,setKeyValueRequest.Value)

	setKeyValueResponse := contract.SetKeyValueResponse{
		Key: setKeyValueRequest.Key,
		Value: setKeyValueRequest.Value,
	}
	//handle err
	resp, _ := json.Marshal(setKeyValueResponse)

	//how to say content type
	w.Write(resp)
}

func (h *KeyValueHandler) GetValue(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//add logs
	//input validation ? - not really required
	//parse req body
	//call repo
	//return response
	values :=  r.URL.Query()
	keyValues, found := values["key"]
	if !found {
		//throw err
	}
	if len(keyValues) > 1 {
		//too many keys //Bad Request
	}

	value := h.InMemoryRepository.GetValue(ctx,keyValues[0])

	getValueResponse := contract.GetValueResponse{
		Key: keyValues[0],
		Value: value,
	}
	//handle err
	resp, _ := json.Marshal(getValueResponse)

	//how to say content type
	w.Write(resp)
}