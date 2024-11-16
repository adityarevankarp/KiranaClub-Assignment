package main

import (
	"log"
	"net/http"
	"server/handlers"
	"server/services"

	"github.com/gorilla/mux"
)

func main() {
	// Load store data
	services.InitStoreData("store.json")

	// Create the router
	router := mux.NewRouter()

	// Register endpoints
	router.HandleFunc("/api/submit/", handlers.SubmitJob).Methods("POST")
	router.HandleFunc("/api/status", handlers.GetJobStatus).Methods("GET")

	// Start the server
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
