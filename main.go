package main

import (
    "fmt"
    "net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
    _, _ = fmt.Fprintf(w, "This is the Home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
    _, _ = fmt.Fprintf(w, "This is the about page\n")
    _, _ = fmt.Fprintf(w, fmt.Sprintf("BTW, 2 + 2 is %d", addTwoValues(2, 2)))
}

// addTwoValues returns sum of two integers
func addTwoValues(x, y int) int {
    return x + y
}

// main is the entry point of the app
func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/about", About)

    http.ListenAndServe(portNumber, nil)
}
