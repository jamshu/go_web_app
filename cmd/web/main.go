package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jamshu/go_web_app/pkg/config"
	"github.com/jamshu/go_web_app/pkg/handlers"
	"github.com/jamshu/go_web_app/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {

	var app config.AppConfig
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
