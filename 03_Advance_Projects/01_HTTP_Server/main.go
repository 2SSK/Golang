package main

import (
	"go-server/db"
	"go-server/routes"
	"go-server/utils"
	"net/http"
)

// Import routes Handlers
var (
	Welcome       = routes.Welcome
	Quotes        = db.Quotes
	QuotesHandler = routes.QuotesHandler
)

func main() {
	PORT := "8080"
	// Create a new ServeMux
	mux := http.NewServeMux()

	mux.HandleFunc("/", Welcome)
	mux.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		QuotesHandler(w, r, Quotes)
	})

	utils.ListenAndServe(PORT, mux)
}
