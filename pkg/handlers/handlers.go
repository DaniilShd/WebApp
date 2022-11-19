package handlers

import (
	"net/http"

	"github.com/DaniilShd/WebApp/pkg/config"
	"github.com/DaniilShd/WebApp/pkg/render"
)

// Repo the repository used by the hendlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perfom some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Daniil"
	// send
	render.RenderTemplate(w, "about.page.html", &TemplateData{
		StringMap: stringMap,
	})
}
