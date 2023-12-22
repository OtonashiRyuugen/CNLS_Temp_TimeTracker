package main

import (
	"log"
	"net/http"
)

func main() {
	// Define the directory to serve.
	fs := http.FileServer(http.Dir("www/"))

	// Handle all requests by serving a file of the same name from ./www.
	http.Handle("/", fs)

	// Define the port to listen on.
	port := "8080"
	log.Printf("Listening on :%s...", port)

	// Start the server.
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
