package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DaniilShd/WebApp/pkg/config"
	"github.com/DaniilShd/WebApp/pkg/handlers"
	"github.com/DaniilShd/WebApp/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	handlers.NewHandlers(handlers.NewRepository(&app))

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	http.Handle("/ModernWebApplication/", http.StripPrefix("/ModernWebApplication/", http.FileServer(http.Dir("/ModernWebApplication/"))))

	fmt.Println(fmt.Sprintf("Starting%s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
