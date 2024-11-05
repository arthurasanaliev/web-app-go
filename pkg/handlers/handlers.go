package handlers

import (
    "fmt"
    "net/http"
    "html/template"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about.page.tmpl")
}

// renderTemplate renders go-template files
func renderTemplate(w http.ResponseWriter, tmpl string) {
    parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
    err := parsedTemplate.Execute(w, nil)
    if err != nil {
        fmt.Println("error parsing template:", err)
        return
    }
}
