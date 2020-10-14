package main

import (
	"log"
	"net/http"
	"os"
	"v1/core"
	"v1/env"
	"golang-user-microservice/config"

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
	router.HandleFunc("/api/users", core.OnSignup).Methods("POST")
	router.HandleFunc("/api/users", core.OnSignup).Methods("POST")
	router.HandleFunc("/api/users", core.OnGetUsers).Methods("GET")
	router.HandleFunc("/api/users/{uuid}", core.OnDeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/{uuid}", core.OnGetUser).Methods("GET")
	router.HandleFunc("/api/users/{uuid}", core.OnUpdateUser).Methods("PUT")

	handler := cor.Handler(router)
	http.ListenAndServe(config.App.Server.Port, handler)
	log.Printf("Port:%s Listening!", config.App.Server.Port)

}
