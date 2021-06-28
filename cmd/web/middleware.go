package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)


func Nosurf(Next http.Handler) http.Handler {
	csrfHandler:= nosurf.New(Next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(Next http.Handler) http.Handler  {
	return session.LoadAndSave(Next)
}