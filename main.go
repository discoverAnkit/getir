package main

import (
	"github.com/discoverAnkit/getir/handler"
	"github.com/discoverAnkit/getir/repository"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
)

func main() {

	keyValueHandler := handler.KeyValueHandler{
		InMemoryRepository: repository.NewInMemoryClient(cache.New(cache.NoExpiration, cache.NoExpiration)),
	}

	handler.HandleRequests(keyValueHandler)

	log.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":10000", nil))
}