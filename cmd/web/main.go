package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jamshu/go_web_app/pkg/config"
	"github.com/jamshu/go_web_app/pkg/handlers"
	"github.com/jamshu/go_web_app/pkg/render"
)

const portNumber = ":8080"
var session *scs.SessionManager
var app config.AppConfig
// main is the main function
func main() {

	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = true

	render.NewTemplate(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/contact", handlers.Repo.Contact)

	fmt.Printf(fmt.Sprintf("Staring application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
