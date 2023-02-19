package main

import (
	"fmt"
	"net/http"
	"github.com/jamshu/go_web_app/pkg/handlers"
)

const portNumber = ":8080"
// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/contact", handlers.Contact)

	fmt.Printf(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
