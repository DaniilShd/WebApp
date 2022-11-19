package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/DaniilShd/WebApp/pkg/config"
	"github.com/DaniilShd/WebApp/pkg/handlers"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// sets the config for the template
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *handlers.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get the template cache from the appconfig

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("myerr TC")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}

	// err = t.Execute(w, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.html")
	if err != nil {
		log.Fatal(err)
		return myCache, err
	}
	fmt.Println(pages)

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Println(name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Fatal(err)
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			log.Fatal(err)
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				log.Fatal(err)
				return myCache, err
			}

		}
		myCache[name] = ts
	}
	return myCache, nil
}
