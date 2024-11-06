package render

import (
    "html/template"
    "path/filepath"
)

// RenderTemplate gets template cache
func GetTemplateCache() (map[string]*template.Template, error) {
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
