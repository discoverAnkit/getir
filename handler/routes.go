package handler

import "net/http"

func HandleRequests(keyValueHandler KeyValueHandler) {
	//How do you say its a POST?
	http.HandleFunc("/setKeyValue", keyValueHandler.SetKeyValue)

	//How do you say its a GET with a query param?
	http.HandleFunc("/getValue", keyValueHandler.GetValue)
}