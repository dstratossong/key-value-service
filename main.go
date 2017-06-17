package main

import (
	"github.com/dstratossong/key-value-service/model"
	"log"
	"net/http"
	"time"
)

const (
	GET     = http.MethodGet
	POST    = http.MethodPost
	PATCH   = http.MethodPatch
	DELETE  = http.MethodDelete
	OPTIONS = http.MethodOptions
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

var Routes = []Route{
	Route{
		Method:  GET,
		Pattern: "/services",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, model.GetServices())
		},
	},
	Route{
		Method:  POST,
		Pattern: "/services/register",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			service := new(model.Service)
			readJSON(r, service)
			model.RegisterService(service)
		},
	},
	Route{
		Method:  GET,
		Pattern: "/fetch/{id}/{key}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
		},
	},
}

func main() {
	server := &http.Server{
		Handler:      MakeRouter(),
		Addr:         ":8080",
		WriteTimeout: 300 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(server.ListenAndServe())
}
