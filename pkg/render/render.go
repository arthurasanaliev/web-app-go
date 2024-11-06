package render

import (
    "fmt"
    "net/http"
    "html/template"
)

// RenderTemplateOnline renders go-template files online (w/o cache)
func RenderTemplateOnline(w http.ResponseWriter, tmpl string) {
    parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
    err := parsedTemplate.Execute(w, nil)
    if err != nil {
        fmt.Println("error parsing template:", err)
        return
    }
}

// TODO -- add cache for templates

// RenderTemplate renders go-templates files if not found in cache
func RenderTemplate(w http.ResponseWriter, tmpl string) {

}
