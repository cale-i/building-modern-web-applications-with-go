package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cale-i/building-modern-web-applications-with-go/pkg/config"
	"github.com/cale-i/building-modern-web-applications-with-go/pkg/handlers"
	"github.com/cale-i/building-modern-web-applications-with-go/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tempCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tempCache
	app.UseCache = false

	repo := handlers.NewPepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
