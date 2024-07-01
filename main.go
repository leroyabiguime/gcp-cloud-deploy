package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World from Cloudor !!")
}

func main() {
	// Define a route for the root URL ("/") that calls the helloWorldHandler function
	http.HandleFunc("/", helloWorldHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
