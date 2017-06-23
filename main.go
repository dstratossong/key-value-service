package main

import (
	"github.com/dstratossong/key-value-service/model"
	"log"
	"net/http"
	"time"
)

var Endpoints = []Endpoint{
	Endpoint{
		Method:  GET,
		Pattern: "/status",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, "Available")
		},
	},
	Endpoint{
		Method:  GET,
		Pattern: "/services",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, model.GetServices())
		},
	},
	Endpoint{
		Method:  POST,
		Pattern: "/services/register",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			service := new(model.Service)
			readJSON(r, service)
			model.RegisterService(service)
		},
	},
	Endpoint{
		Method:  GET,
		Pattern: "/run/fetch/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, model.GetPropertyForRun(getUint64Param(r, "id")))
		},
	},
	Endpoint{
		Method:  GET,
		Pattern: "/run/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, model.GetRun(getUint64Param(r, "id")))
		},
	},
	Endpoint{
		Method:  POST,
		Pattern: "/run/create",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			obj := new(struct {
				Service    string
				Properties *model.Properties
			})
			readJSON(r, obj)
			writeJSON(w, model.NewServiceRun(obj.Service, obj.Properties))
		},
	},
	Endpoint{
		Method:  POST,
		Pattern: "/run/finish/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			id := getUint64Param(r, "id")
			obj := new(struct {
				Status string
				Data   *model.Properties
			})
			readJSON(r, obj)
			model.FinishRun(id, obj.Status, obj.Data)
		},
	},
}

func main() {
	server := &http.Server{
		Handler:      MakeRouter(Endpoints),
		Addr:         ":8080",
		WriteTimeout: 300 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(server.ListenAndServe())
}
