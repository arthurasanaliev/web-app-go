package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// AppConfig is app-config data container
type AppConfig struct {
	InProduction bool
	UseCache     bool
	TempCache    map[string]*template.Template
	Session      *scs.SessionManager
}
