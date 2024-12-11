package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve static files from the "frontend" directory
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	// Simple API endpoint
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "Hello from Golang API!"}`)
	})

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
