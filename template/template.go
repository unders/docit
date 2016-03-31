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

// Data content visible on the HTML page.
type Data struct {
	Name string
}

var page Page
var data Data

// Init initialize template
func Init(box *rice.Box, d Data) {
	data = d

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
func Render(w http.ResponseWriter, b []byte, statusCode int) {
	body := struct {
		Body tmpl.HTML
		Name string
	}{
		Body: tmpl.HTML(b),
		Name: data.Name,
	}

	w.WriteHeader(statusCode)

	if err := page.Tmpl.Execute(w, body); err != nil {
		log.Printf("template.Render(w, makdown, %d). Error: %v", statusCode, err)
		http.Error(w, "500: Internal Server Error", http.StatusInternalServerError)
	}
}
