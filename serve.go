package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/russross/blackfriday"
	"github.com/unders/docit/cli"
	"github.com/unders/docit/template"
)

const notFoundMsg = "<pre style='word-wrap: break-word;" +
	"white-space: pre-wrap;'>404 page not found </pre>"

func loadPage(filename string) ([]byte, int) {
	markdown, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Markdown file %s not found. Error: %v\n", filename, err)
		return []byte(notFoundMsg), http.StatusNotFound
	}

	html := blackfriday.MarkdownCommon(markdown)

	return html, http.StatusOK
}

func serve(arg cli.Arg) {
	fileServer := http.FileServer(http.Dir(arg.Root))

	root := func(w http.ResponseWriter, req *http.Request) {
		upath := req.URL.Path
		if !strings.HasPrefix(upath, "/") {
			upath = "/" + upath
			req.URL.Path = upath
		}

		// Redirect to index.md page.
		if req.URL.Path == "/" {
			http.Redirect(w, req, "/"+arg.Index, http.StatusSeeOther)
			return
		}

		// Redirect to help.md page.
		if req.URL.Path == "/help" {
			http.Redirect(w, req, "/"+arg.Help, http.StatusSeeOther)
			return
		}

		// Parse and serve given Markdown file relative to root dir.
		if strings.HasSuffix(req.URL.Path, ".md") {
			html, code := loadPage(arg.Root + req.URL.Path)
			template.Render(w, html, code)
			return
		}

		// Serve file relative to root dir.
		fileServer.ServeHTTP(w, req)
	}
	http.HandleFunc("/", root)

	data := template.Data{
		Name: arg.Name,
	}
	template.Init(rice.MustFindBox("embedded_assets/tmpl"), data)

	box := rice.MustFindBox("embedded_assets")
	embeddedFileServer := http.StripPrefix("/embedded_assets/", http.FileServer(box.HTTPBox()))
	http.Handle("/embedded_assets/", embeddedFileServer)

	fmt.Printf("Serving static files at http://0.0.0.0:%s from dir %s\n", arg.Port, arg.Root)
	log.Fatal(http.ListenAndServe(":"+arg.Port, nil))
}
