package root

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
	"github.com/unders/docit/cli"
	"github.com/unders/docit/template"
)

const notFoundMsg = "<pre style='word-wrap: break-word;" +
	"white-space: pre-wrap;'>404 page not found</pre>"

// Handle renders pages unders "/"
func Handle(arg cli.Arg) func(w http.ResponseWriter, req *http.Request) {
	fileServer := http.FileServer(http.Dir(arg.Root))

	return func(w http.ResponseWriter, req *http.Request) {
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
			docs := listDoc(arg.Root, req.URL.Path)
			template.Render(w, html, docs, code)
			return
		}

		// Serve file relative to root dir.
		fileServer.ServeHTTP(w, req)
	}
}

func listDoc(root, p string) []template.Doc {
	base, _ := path.Split(p)
	dir := root + base

	if !strings.Contains(dir, "doc/") {
		if _, err := os.Stat(dir + "doc/"); err != nil {
			return []template.Doc{}
		}
		dir = dir + "doc/"
	}

	var docs []template.Doc
	walk := func(filepath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		link := filepath[len(root):]

		name := ""
		if i := strings.LastIndex(link, "doc/"); i != -1 {
			name = link[i:]
		} else {
			_, name = path.Split(filepath)
		}

		doc := template.Doc{
			Link: link,
			Name: name,
		}
		docs = append(docs, doc)
		return err
	}

	if err := filepath.Walk(dir, walk); err != nil {
		log.Println("walk dir error ", err)
		return []template.Doc{}
	}

	return docs
}

func loadPage(filename string) ([]byte, int) {
	markdown, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Markdown file %s not found. Error: %v\n", filename, err)
		return []byte(notFoundMsg), http.StatusNotFound
	}

	html := blackfriday.MarkdownCommon(markdown)

	return html, http.StatusOK
}
