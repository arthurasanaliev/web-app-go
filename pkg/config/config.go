package config

import "html/template"

// AppConfig is app-config data container
type AppConfig struct {
    TempCache map[string]*template.Template
}
