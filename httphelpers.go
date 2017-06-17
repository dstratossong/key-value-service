package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Writes CORS headers to the response
func writeCORSHeader(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)      // SSL requires matching origin, * does not work with SSL
	w.Header().Set("Access-Control-Allow-Credentials", "true") // enable SSL
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Wrapper for normal logging requests
// For internal endpoints, only outputs to log file
func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(">>", r.Method, r.RequestURI)
		start := time.Now()

		inner.ServeHTTP(w, r)

		finish := time.Now()
		log.Println("<<", "Served in", finish.Sub(start), "\n")
	})
}

// Wrapper for CORS OPTIONS request
// Writes a CORS header to requests and handle the OPTIONS method
func HandleOptions(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			writeCORSHeader(w, r)
			return
		} else {
			writeCORSHeader(w, r)
			h.ServeHTTP(w, r)
		}
	}
}

// Returns a fully configured router based on Routes
func MakeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range Routes {
		router.Methods(route.Method, OPTIONS).Path(route.Pattern).Handler(Logger(HandleOptions(route.Handler)))
	}
	return router
}

// Reads JSON from a HTTP Request and marshal it into an obj
func readJSON(r *http.Request, obj interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(obj)
	if err != nil {
		panic(err)
	}

	log.Printf("Read JSON request body:\n%+v\n", obj)
}

// Writes JSON to a HTTP Response with Status 200 OK
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
