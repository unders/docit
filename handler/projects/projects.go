package projects

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/unders/docit/template"
)

const notFoundMsg = "<pre style='word-wrap: break-word;" +
	"white-space: pre-wrap;'>404 No projects found</pre>"

const dirErrMsg = "<pre style='word-wrap: break-word;" +
	"white-space: pre-wrap;'>Projects dir could not be read</pre>"

func isProject(root string, f os.FileInfo) bool {
	if f.IsDir() {
		return true
	}

	if f.Mode()&os.ModeSymlink == 0 {
		return false
	}

	dir, err := os.Readlink(root + "/" + f.Name())
	if err != nil {
		return false
	}

	if !filepath.IsAbs(dir) {
		dir = root + "/" + dir
	}

	fileInfo, err := os.Stat(dir)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

// Handle renders projects in root directory.
func Handle(root string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		dir, err := ioutil.ReadDir(root)

		if err != nil {
			log.Printf("Could not read root directory.\n")
			template.Render(w, []byte(dirErrMsg), []template.Doc{},
				http.StatusInternalServerError)
		}

		length := len(dir)

		if length == 0 {
			log.Printf("No directory found, len(dir) == %d\n", length)
			template.Render(w, []byte(notFoundMsg), []template.Doc{},
				http.StatusNotFound)
		}

		projects := make([]template.Project, length)

		for i, file := range dir {
			if isProject(root, file) {
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
