package main

import (
	"golang-user-microservice/config"
	"golang-user-microservice/core"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind address for the SSV verify server")

func main() {
	// err := env.Parse()
	// if err != nil {
	// 	log.Println("Error parsing env", "error", err)
	// 	os.Exit(1)
	// }

	config.ReadConfig()

	cor := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin", "Content-Type", "Session-Key", "Device-ID"},
		Debug:          true,
	})

	router := mux.NewRouter()
	router.HandleFunc("/api/ping", core.OnPing).Methods("GET")

	handler := cor.Handler(router)
	http.ListenAndServe(config.GetPort(), handler)
	log.Println("Listening!")
}
