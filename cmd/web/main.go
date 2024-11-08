package main

import (
	"fmt"
    "github.com/alexedwards/scs/v2"
	"github.com/arthurasanaliev/web-app-go/pkg/config"
	"github.com/arthurasanaliev/web-app-go/pkg/handlers"
	"github.com/arthurasanaliev/web-app-go/pkg/render"
    "time"
	"log"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the entry point of the app
func main() {
    app.InProduction = false

    session = scs.New()
    session.Lifetime = 24 * time.Hour
    session.Cookie.Persist = true
    session.Cookie.SameSite = http.SameSiteLaxMode
    session.Cookie.Secure = app.InProduction 

    app.Session = session

	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not creat template cache")
	}

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
