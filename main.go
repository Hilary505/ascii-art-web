package main

import (
	
	"fmt"
	"log"
	"net/http"
	"ascii/handlers"
)

// main starts the HTTP server.
func main() {
	fmt.Println("Server is starting...")
	fmt.Println("Server is running on http://localhost:8080")
	http.HandleFunc("/", handlers.Handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
