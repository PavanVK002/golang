package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // Registering the helloHandler function to handle requests to the /hello route
    http.HandleFunc("/hello", helloHandler)

    // Start the HTTP server on port 8080
    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}
