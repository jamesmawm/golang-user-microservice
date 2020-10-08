package main

import (
	"golang-user-microservice/config"
	"golang-user-microservice/core"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	config.ReadConfig()

	cor := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin", "Content-Type", "Session-Key", "Device-ID"},
		Debug:          true,
	})

	router := mux.NewRouter()
	router.HandleFunc("/api/ping", core.OnPing).Methods("GET")

	handler := cor.Handler(router)
	http.ListenAndServe(config.App.Server.Port, handler)
	log.Printf("Port:%s Listening!", config.App.Server.Port)

}
