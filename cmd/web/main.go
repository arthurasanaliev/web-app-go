package main

import (
	"fmt"
	"github.com/arthurasanaliev/web-app-go/pkg/config"
	"github.com/arthurasanaliev/web-app-go/pkg/handlers"
	"github.com/arthurasanaliev/web-app-go/pkg/render"
	"log"
	"net/http"
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

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Starting application on port", portNumber)

	_ = srv.ListenAndServe()
}
