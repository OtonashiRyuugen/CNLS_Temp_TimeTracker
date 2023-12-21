package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

func main() {
	// Specify the domain for which you want to get a TLS certificate
	domain := "cnls.io"

	// Setup automatic certificate management with Let's Encrypt via Cloudflare DNS
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache("certs"), // Store certificates in a cache directory
	}

	// Create a new HTTP server
	server := &http.Server{
		Addr:    "[::]:443", // Listen on all available addresses for IPv4 and IPv6
		Handler: http.HandlerFunc(handler),
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
		},
	}

	// Print a message indicating the server is running
	fmt.Printf("Server is running on https://%s\n", domain)

	// Start the HTTPS server
	err := server.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Your HTML content here
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Go Web Server</title>
		</head>
		<body>
			<h1>Hello, this is your Go web server!</h1>
		</body>
		</html>
	`

	// Write the HTML content to the response writer
	w.Write([]byte(html))
}
