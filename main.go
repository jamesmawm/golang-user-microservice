package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"v1/core"
)

var defaultPort = ":8080"

func main() {
	cor := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin", "Content-Type", "Session-Key", "Device-ID"},
		Debug:          true,
	})

	router := mux.NewRouter()
	router.HandleFunc("/api/ping", core.OnPing).Methods("GET")

	port := defaultPort
	handler := cor.Handler(router)
	http.ListenAndServe(port, handler)
	log.Println("Listening!")
}
