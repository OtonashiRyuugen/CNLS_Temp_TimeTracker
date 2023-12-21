package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/crypto/acme/autocert"
	"net"
	"net/http"
)

func main() {
	// Specify the domain for which you want to get a TLS certificate
	domain := "your-domain.com"

	// Setup automatic certificate management with Let's Encrypt via Cloudflare DNS
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache("certs"), // Store certificates in a cache directory
	}

	// Create a TCP listener on a specific IPv4 address and port 443
	listener, err := net.Listen("tcp4", ":443")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a new HTTP server
	server := &http.Server{
		Addr:    ":443",
		Handler: http.HandlerFunc(handler),
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
		},
	}

	// Print a message indicating the server is running
	fmt.Printf("Server is running on https://%s\n", domain)

	// Start the HTTPS server
	err = server.ServeTLS(listener, "", "")
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
