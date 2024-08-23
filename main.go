package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the home page!")
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Create a new http.Server
	server := &http.Server{
		Addr:    ":8080", // Set the address to listen on port 8080
		Handler: mux,     // Use the created ServeMux as the handler
	}

	// Start the server
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed:", err)
	}
}
