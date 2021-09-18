package handler

import "net/http"

func HandleRequests(keyValueHandler KeyValueHandler) {
	//How do you say its a POST?
	http.HandleFunc("/setKeyValue", keyValueHandler.setKeyValue)
}