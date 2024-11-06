package main

import (
    "fmt"
    "net/http"
    "github.com/arthurasanaliev/web-app-go/pkg/handlers"
    "github.com/arthurasanaliev/web-app-go/pkg/render"
)

const portNumber = ":8080"

// main is the entry point of the app
func main() {
    tmplCache, err := render.GetTemplateCache()
    if err != nil {
        fmt.Println("cannot cache templates")
        return
    }

    handlers.CreateMap(tmplCache)

    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.About)

    _ = http.ListenAndServe(portNumber, nil)
}
