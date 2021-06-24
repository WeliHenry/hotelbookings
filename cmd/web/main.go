package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/welihenry/hotelbookings/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"



func main()  {
	//session := scs.New()
	//session.Lifetime = 24 * time.Hour
	//session.Cookie.Persist = true
	//session.Cookie.SameSite = http.SameSiteLaxMode
	//session.Cookie.Secure = false

	router:= chi.NewRouter()
	router.Get("/", handlers.Home)
	log.Fatal(http.ListenAndServe(portNumber, router))





}

