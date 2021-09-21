package main

import (
	"context"
	"github.com/discoverAnkit/getir/handler"
	"github.com/discoverAnkit/getir/repository"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"os"
)

// Connection URI
// It can come from a config service(a separate config package can be created then to deal with that service) as well if there is one
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

	mongoRequestHandler := &handler.MongoRequestHandler{MongoRepo: mongoStore}
	keyValueHandler := &handler.KeyValueHandler{
		InMemoryRepository: repository.NewInMemoryClient(cache.New(cache.NoExpiration, cache.NoExpiration)),
	}

	router :=  handler.NewRouter(keyValueHandler,mongoRequestHandler)
	http.HandleFunc("/", router.Handle)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "1000"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Starting http server on port : ",port)
}