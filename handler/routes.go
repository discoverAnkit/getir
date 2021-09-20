package handler

import (
	"context"
	"net/http"
)

type Router struct {
	keyValueHandler     *KeyValueHandler
	mongoRequestHandler *MongoRequestHandler
}

func NewRouter(keyValueHandler *KeyValueHandler, mongoRequestHandler *MongoRequestHandler) *Router {
	return &Router{
		keyValueHandler: keyValueHandler,
		mongoRequestHandler: mongoRequestHandler,
	}
}

func (s *Router) Handle (w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		case "/getValue":
			s.keyValueHandler.GetValue(ctx,w,r)
		default:
			http.NotFound(w,r)
		}
	case "POST":
		switch r.URL.Path {
		case "/setKeyValue":
			s.keyValueHandler.SetKeyValue(ctx,w,r)
		case "/getKeyValueRecords":
			s.mongoRequestHandler.GetKeyValueRecords(ctx,w,r)
		default:
			http.NotFound(w,r)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}