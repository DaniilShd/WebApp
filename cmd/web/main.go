package main

import (
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

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// fmt.Println(fmt.Sprintf("Starting%s", portNumber))

	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
