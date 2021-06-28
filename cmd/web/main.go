package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/welihenry/hotelbookings/pkg/config"
	"github.com/welihenry/hotelbookings/pkg/handlers"
	"github.com/welihenry/hotelbookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {


	session = scs.New()
	session.Lifetime = 24 *time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session







	app.InProduction = false
	tc, err:= render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo:= handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)


	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(Nosurf)
	router.Use(SessionLoad)

	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	log.Fatal(http.ListenAndServe(portNumber, router))

}
