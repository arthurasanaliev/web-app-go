package handlers

import (
    "net/http"
    "github.com/arthurasanaliev/web-app-go/pkg/render"
    "github.com/arthurasanaliev/web-app-go/pkg/config"
)

var Repo *Repository

// Repository holds repo info for handlers
type Repository struct {
    App *config.AppConfig
}

// NewRepo creates new Repository instance
func NewRepo(a *config.AppConfig) *Repository {
    return &Repository{
        App: a,
    }
}

// SetRepo sets Repository
func SetRepo(r *Repository) {
    Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "about.page.tmpl")
}
