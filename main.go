package main

import (
	"asciiart/serv"  
	"log"            
	"net/http"      
)

func main() {
	// Handle the root URL path ("/") by mapping it to the Index function in the serv package
	// This will serve as the homepage for the application
	http.HandleFunc("/", serv.Index)

	// Handle the "/ascii-art" URL path by mapping it to the AsciiWeb function in the serv package
	// This endpoint will likely handle the ASCII art generation and serve the related output
	http.HandleFunc("/ascii-art", serv.AsciiWeb)

	// Log a message to the console to inform that the server is running and accessible at localhost:8080
	log.Println("Server running at http://localhost:8080")

	// Start the HTTP server on port 8080 and listen for incoming requests
	// If the server encounters an error, ListenAndServe will return it
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err) // Log fatal error if the server fails
	}
}
