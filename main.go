package main

import (
	"fmt"
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

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}