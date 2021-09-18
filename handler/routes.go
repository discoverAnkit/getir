package handler

import (
	"context"
	"net/http"
)

func HandleRequests(keyValueHandler KeyValueHandler) {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()

		switch r.Method {
		case "GET":
			switch r.URL.Path {
			case "/getValue":
				keyValueHandler.GetValue(ctx,w,r)
			default:
				http.NotFound(w,r)
			}
		case "POST":
			switch r.URL.Path {
			case "/setKeyValue":
				keyValueHandler.SetKeyValue(ctx,w,r)
			default:
				http.NotFound(w,r)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}