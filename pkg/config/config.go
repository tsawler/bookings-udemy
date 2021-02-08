package config

import "html/template"

// AppConfig holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
