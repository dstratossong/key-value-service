package main

import (
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
		Method:  POST,
		Pattern: "/fetch",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, "Hello world!")
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
