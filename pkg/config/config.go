package config

import "html/template"

// AppConfig is app-config data container
type AppConfig struct {
    UseCache bool
    TempCache map[string]*template.Template
}
