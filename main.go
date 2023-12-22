package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve files from the www directory.
	fs := http.FileServer(http.Dir("www/"))

	// Map the root URL to the www directory.
	http.Handle("/", fs)

	// Define the port to listen on.
	port := "8080" // You can change this to 80 for standard HTTP
	log.Printf("Server is listening on port %s...", port)

	// Start the server.
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
