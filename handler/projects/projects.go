package projects

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/unders/docit/template"
)

const notFoundMsg = "<pre style='word-wrap: break-word;" +
	"white-space: pre-wrap;'>404 No projects found</pre>"

const dirErrMsg = "<pre style='word-wrap: break-word;" +
	"white-space: pre-wrap;'>Projects dir could not be read</pre>"

// Handle renders projects in root directory.
func Handle(root string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		dir, err := ioutil.ReadDir(root)

		if err != nil {
			log.Printf("Could not read root directory.\n")
			template.Render(w, []byte(dirErrMsg),
				http.StatusInternalServerError)
		}

		length := len(dir)

		if length == 0 {
			log.Printf("No directory found, len(dir) == %d\n", length)
			template.Render(w, []byte(notFoundMsg),
				http.StatusNotFound)
		}

		projects := make([]template.Project, length)

		for i, file := range dir {
			if file.IsDir() {
				name := file.Name()

				proj := template.Project{
					Link: strings.Join([]string{"/", name, "/Readme.md"}, ""),
					Name: name,
				}
				projects[i] = proj
			}
		}
		template.RenderProjects(w, projects)
	}
}
