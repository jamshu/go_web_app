package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl,"./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

var tc = make(map[string]*template.Template)
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	
	
	_,isData := tc[t]
	if !isData {
		//need to create template
		log.Println("Creating Cache for the tempate")
		err = createTemplate(t)
		if err != nil {
			fmt.Println("Error in creting  template cache:",err)
		}
	} else {
		log.Println("Using Cached template")
		
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error in parsing template:",err)
	}

}
func createTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s",t),
		"./templates/base.layout.tmpl",	
	}
	log.Println("templates:", templates)
	tmpl,err := template.ParseFiles(templates...)
	if err != nil {
		return err 
	}
	tc[t] = tmpl
	return nil
}