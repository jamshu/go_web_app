package handlers

import (
	"net/http"
	"github.com/jamshu/go_web_app/pkg/render"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

// Contact is the handler for the contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.html")
}


