package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Welcome to My Go Web Page</title>
			</head>
			<body>
				<h1>Hello from Go!</h1>
				<p>This is a basic HTML page served by a Go program.</p>
			</body>
			</html>
		`
		fmt.Fprint(w, html)
	})

	port := 8080
	fmt.Printf("Starting server on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
