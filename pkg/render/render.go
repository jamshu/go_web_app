package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/jamshu/go_web_app/pkg/config"
	"github.com/jamshu/go_web_app/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData (data *models.TemplateData) *models.TemplateData {
	return data
}
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	// create templates cache
	// tc,err := CreateTemplateCache()
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Println("Not found template")
	}
	data = AddDefaultData(data)
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}

}
func CreateTemplateCache() (map[string]*template.Template, error) {
	//create template cache
	myCache := map[string]*template.Template{}
	//get all the files to create the cache

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//loop through pages

	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//get the layout file from disk
		layout, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(layout) > 0 {
			ts.ParseGlob("./templates/*.layout.tmpl")
		}
		myCache[name] = ts

	}
	return myCache, nil

}
