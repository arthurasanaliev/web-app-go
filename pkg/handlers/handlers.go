package handlers

import (
    "path/filepath"
    "net/http"
    "html/template"
)

var tmplCache map[string]*template.Template

// CreateMap sets template cache
func CreateMap(cache map[string]*template.Template) {
    tmplCache = cache
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
    // TODO -- create render page function
    name := filepath.Base("./templates/home.page.tmpl")
    tmplCache[name].Execute(w, nil)
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
    // TODO -- create render page function
    name := filepath.Base("./templates/about.page.tmpl")
    tmplCache[name].Execute(w, nil)
}
