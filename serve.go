package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/russross/blackfriday"
	"github.com/unders/docit/cli"
	"github.com/unders/docit/template"
)

const notFoundMsg = "<pre style='word-wrap: break-word;" +
	"white-space: pre-wrap;'>404 page not found</pre>"

const projectNotFoundMsg = "<pre style='word-wrap: break-word;" +
	"white-space: pre-wrap;'>404 No projects found</pre>"

const projectDirErrMsg = "<pre style='word-wrap: break-word;" +
	"white-space: pre-wrap;'>Projects dir could not be read</pre>"

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
	http.HandleFunc("/projects", projects(arg))
	if arg.MemberFile != "" {
		path := arg.Root + "/" + arg.MemberFile
		http.HandleFunc("/members", members(path))
	}

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

func projects(arg cli.Arg) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		dir, err := ioutil.ReadDir(arg.Root)

		if err != nil {
			log.Printf("Could not read root directory.\n")
			template.Render(w, []byte(projectDirErrMsg),
				http.StatusInternalServerError)

		}

		length := len(dir)

		if length == 0 {
			log.Printf("No directory found, len(dir) == %d\n", length)
			template.Render(w, []byte(projectNotFoundMsg),
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

func members(path string) func(w http.ResponseWriter, req *http.Request) {
	m := readMembers(path)

	length := len(m)

	members := make([]template.Member, length)

	for i, mm := range m {
		members[i].Email = mm
	}

	return func(w http.ResponseWriter, req *http.Request) {
		template.RenderMembers(w, members)
	}
}

func readMembers(path string) []string {
	var members []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file %s, err %v", path, err)
	}
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		members = append(members, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Could not scan file %s, err %v", path, err)
	}

	return members
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatalf("Could not close file: %v, %v", f, err)
	}
}
