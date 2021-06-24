package handlers

import (
	"github.com/welihenry/hotelbookings/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request)  {

	render.RenderTemplate(w, "home.html")

}

func About(w http.ResponseWriter, r *http.Request)  {

}


