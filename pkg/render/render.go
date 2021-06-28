package render

import (
	"bytes"
	"fmt"
	"github.com/welihenry/hotelbookings/pkg/config"
	"github.com/welihenry/hotelbookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func AddDefaultTemplateData(data *models.TemplateData) *models.TemplateData {

	return data
}
func NewTemplates(a *config.AppConfig)  {
	app = a
}
    var tc map[string]*template.Template
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData)  {

	if app.UseCache {

		tc= app.TemplateCache
	}else {
		tc,_ = CreateTemplateCache()
	}


	t, ok:= tc[tmpl]
	if !ok {
		log.Fatal("could not get template from cache")
	}
	buf:= new(bytes.Buffer)
	_= t.Execute(buf, data)
	_,err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

	parsedTemplate,_:= template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing the template", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err:= filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name:= filepath.Base(page)
		ts, err:= template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache,err
		}
		matches, err:= filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache,err
			}
		}
		myCache[name] = ts

	}
	return myCache,nil
}
