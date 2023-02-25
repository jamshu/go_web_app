package handlers

import (
	"net/http"

	"github.com/jamshu/go_web_app/pkg/config"
	"github.com/jamshu/go_web_app/pkg/models"
	"github.com/jamshu/go_web_app/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello world!"
	remoteIP := r.RemoteAddr
	my_secret_key := "its my secret key 9880000"
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	m.App.Session.Put(r.Context(), "secret_key", my_secret_key)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	secret_key := m.App.Session.GetString(r.Context(),"secret_key")
	stringMap["remote_ip"] = remoteIP
	stringMap["secret_key"] = secret_key
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact is the handler for the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
