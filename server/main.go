package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jtoppings/volleyball-server/internal/api"
)

func main() {
	router := mux.NewRouter()

	// Define your API routes
	router.HandleFunc("/api/listResponsibilitiesQuestions", api.ListResponsibilitiesQuestionsHandler).Methods("GET")

	http.Handle("/", router)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
