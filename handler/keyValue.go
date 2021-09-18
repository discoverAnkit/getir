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

func (h *KeyValueHandler) setKeyValue(w http.ResponseWriter, r *http.Request) {
	//add logs
	//input validation ? - not really required 
	//parse req body
	//call repo
	//return response
	ctx := context.Background()
	setKeyValueRequest := contract.SetKeyValueRequest{}
	json.NewDecoder(r.Body).Decode(&setKeyValueRequest)
	h.InMemoryRepository.SetKeyValue(ctx,setKeyValueRequest.Key,setKeyValueRequest.Value)

	setKeyValueResponse := contract.SetKeyValueResponse{
		Key: setKeyValueRequest.Key,
		Value: setKeyValueRequest.Value,
	}
	//handle err
	resp, _ := json.Marshal(setKeyValueResponse)

	w.Write(resp)
}
