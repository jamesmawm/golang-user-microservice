package main

import (
	"log"
	"net/http"
	"os"
	"v1/core"

	"github.com/karimelazzouni/golang-user-microservice/env"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind address for the SSV verify server")

func main() {
	err := env.Parse()
	if err != nil {
		log.Println("Error parsing env", "error", err)
		os.Exit(1)
	}
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

	port := *bindAddress
	handler := cor.Handler(router)
	http.ListenAndServe(port, handler)
	log.Println("Listening!")
}
