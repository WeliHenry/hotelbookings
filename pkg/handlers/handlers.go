package handlers

import (
	"github.com/welihenry/hotelbookings/pkg/config"
	"github.com/welihenry/hotelbookings/pkg/models"
	"github.com/welihenry/hotelbookings/pkg/render"
	"net/http"
)



var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig)  *Repository{
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository)  {
	Repo = r
}



func (m *Repository) Home(w http.ResponseWriter, r *http.Request)  {
	RemoteIp:= r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", RemoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})




}

func (m *Repository) About(w http.ResponseWriter, r *http.Request)  {
	StringMap := make(map[string]string)
	StringMap["test"]= "hello world"

	RemoteIp:= m.App.Session.GetString(r.Context(),"remote_ip")
	StringMap["remote_ip"] = RemoteIp


	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: StringMap})

}


