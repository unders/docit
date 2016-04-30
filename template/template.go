package template

import (
	tmpl "html/template"
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
)

// Page is used for rendering html pages.
type Page struct {
	Tmpl *tmpl.Template
}

// Project struct for creating a link to a project.
type Project struct {
	Link string
	Name string
}

// Doc struct for creating a link to a doc page
type Doc struct {
	Link string
	Name string
}

// Member struct for creating a Member page.
type Member struct {
	Email string
}

// Data content visible on the HTML page.
type Data struct {
	Name      string
	Body      tmpl.HTML
	Projects  []Project
	Members   []Member
	Docs      []Doc
	HasProj   bool
	HasMember bool
	HasDoc    bool
}

var page Page
var data Data

// Init initialize template
func Init(box *rice.Box, name string) {
	data.Name = name

	layout, err := box.String("layout.html")
	if err != nil {
		log.Fatal(err)
	}

	page.Tmpl, err = tmpl.New("page").Parse(layout)
	if err != nil {
		log.Fatal(err)
	}
}

// Render a HTML page
func Render(w http.ResponseWriter, b []byte, docs []Doc, statusCode int) {
	hasDoc := len(docs) != 0
	d := Data{
		Body:   tmpl.HTML(b),
		Name:   data.Name,
		Docs:   docs,
		HasDoc: hasDoc,
	}

	render(w, d, statusCode)
}

// RenderProjects render projects
func RenderProjects(w http.ResponseWriter, p []Project) {
	d := Data{
		Name:     data.Name,
		Projects: p,
		HasProj:  true,
	}

	render(w, d, http.StatusOK)
}

// RenderMembers render members
func RenderMembers(w http.ResponseWriter, m []Member) {
	d := Data{
		Name:      data.Name,
		Members:   m,
		HasMember: true,
	}

	render(w, d, http.StatusOK)
}

func render(w http.ResponseWriter, d Data, statusCode int) {
	w.WriteHeader(statusCode)

	if err := page.Tmpl.Execute(w, d); err != nil {
		log.Printf("template.render. status = %d, error = %v", statusCode, err)
		http.Error(w, "500: Internal Server Error", http.StatusInternalServerError)
	}
}
