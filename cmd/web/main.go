package main

import (
    "log"
    "net/http"
    "github.com/arthurasanaliev/web-app-go/pkg/render"
    "github.com/arthurasanaliev/web-app-go/pkg/config"
    "github.com/arthurasanaliev/web-app-go/pkg/handlers"
)

const portNumber = ":8080"

// main is the entry point of the app
func main() {
    tmplCache, err := render.CreateTemplateCache()
    if err != nil {
        log.Fatal("could not creat template cache")
    }

    var app config.AppConfig
    app.TempCache = tmplCache
    app.UseCache = true
    render.SetApp(&app)

    repo := handlers.NewRepo(&app)
    handlers.SetRepo(repo)

    http.HandleFunc("/", handlers.Repo.Home)
    http.HandleFunc("/about", handlers.Repo.About)

    _ = http.ListenAndServe(portNumber, nil)
}
