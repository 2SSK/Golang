package utils

import (
	"log"
	"net/http"
)

func ListenAndServe(port string, handler http.Handler) {
	// Log the message before starting the server
	log.Printf("server started on port %s...\n", port)
	// listen to port
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatalf("server failed to start: %v\n", err)
		return
	}
}
