package main

import (
    "net/http"
    "github.com/arthurasanaliev/web-app-go/pkg/handlers"
)

const portNumber = ":8080"

// main is the entry point of the app
func main() {
    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.About)

    _ = http.ListenAndServe(portNumber, nil)
}
