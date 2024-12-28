package main

import (
	"fmt"
	"net/http"

	"main/database"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/{endpoint}", incoming).Methods("GET")

	handler := handlers.CORS(
		//allow all origins (FOR NOW)
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET"}),
		handlers.AllowedHeaders([]string{"Content-Type", "application/json"}),
	)

	http.Handle("/", handler(r))
	port := "4001"

	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func incoming(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	e := vars["endpoint"]

	switch e {
	case "get_default_list":
		database.GetDefaultList()

	default:
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid endpoint", http.StatusBadGateway)
	}
}
