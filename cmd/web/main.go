package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/welihenry/hotelbookings/pkg/config"
	"github.com/welihenry/hotelbookings/pkg/handlers"
	"github.com/welihenry/hotelbookings/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	//session := scs.New()
	//session.Lifetime = 24 * time.Hour
	//session.Cookie.Persist = true
	//session.Cookie.SameSite = http.SameSiteLaxMode
	//session.Cookie.Secure = false

	var app config.AppConfig
	tc, err:= render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo:= handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)


	router := chi.NewRouter()
	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	log.Fatal(http.ListenAndServe(portNumber, router))

}
