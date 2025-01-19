package routes

import (
	"encoding/json"
	"fmt"
	"go-server/db"
	"net/http"
)

type Quote = db.Quote

// Functions to handle home route
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the go server!")
}

// Function to handle quotes route
func QuotesHandler(w http.ResponseWriter, r *http.Request, data []Quote) {
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the data to JSON and write it to the response writer
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode Quotes", http.StatusInternalServerError)
	}
}
