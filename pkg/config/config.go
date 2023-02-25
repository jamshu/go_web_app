package config

import "html/template"

// AppConfig hold application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
