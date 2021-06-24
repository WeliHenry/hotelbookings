package main

import (
	"github.com/go-chi/chi/v5"
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
	router.Get("/", Home)
	log.Fatal(http.ListenAndServe(portNumber, router))





}

