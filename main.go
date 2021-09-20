package main

import (
	"context"
	"github.com/discoverAnkit/getir/handler"
	"github.com/discoverAnkit/getir/repository"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
)

// Connection URI
// It can come from a config service as well if there is one
const (
	mongoUri = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?authSource=admin&replicaSet=challenge-shard-0&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true"
	mongoDSN = "getir-case-study"
	mongoCollection = "records"
)


func main() {

	mongoStore := repository.NewMongoStore(mongoUri,mongoDSN,mongoCollection)
	defer func() {
		if err := mongoStore.MongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	mongoRequestHandler := handler.MongoRequestHandler{MongoRepo: mongoStore}
	keyValueHandler := handler.KeyValueHandler{
		InMemoryRepository: repository.NewInMemoryClient(cache.New(cache.NoExpiration, cache.NoExpiration)),
	}

	handler.HandleRequests(keyValueHandler,mongoRequestHandler)

	log.Println("Starting server for testing HTTP POST...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}