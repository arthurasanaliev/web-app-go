package render

import (
	"bytes"
	"github.com/arthurasanaliev/web-app-go/pkg/config"
	"github.com/arthurasanaliev/web-app-go/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// SetApp sets app for render package
func SetApp(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	// TODO -- add default template data
	return td
}

// RenderTemplate renders template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tmplCache map[string]*template.Template
	if app.UseCache {
		tmplCache = app.TempCache
	} else {
		tmplCache, _ = CreateTemplateCache()
	}

	t, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal("cannot get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = addDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache creates template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tmplCache, err
	}

	layouts, err := filepath.Glob("./templates/*layout.tmpl")
	if err != nil {
		return tmplCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmplSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}
		if len(layouts) > 0 {
			tmplSet, err = tmplSet.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return tmplCache, err
			}
		}
		tmplCache[name] = tmplSet
	}

	return tmplCache, nil
}
