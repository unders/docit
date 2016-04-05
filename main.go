package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/unders/docit/cli"
	"github.com/unders/docit/handler/members"
	"github.com/unders/docit/handler/projects"
	"github.com/unders/docit/handler/root"
	"github.com/unders/docit/template"
)

func main() {
	switch cmd, arg := cli.Parse(); true {
	case cmd == "serve":
		drawRoutes(arg)
		fmt.Printf("Serving static files at http://0.0.0.0:%s from dir %s\n", arg.Port, arg.Root)
		log.Fatal(http.ListenAndServe(":"+arg.Port, nil))
	case cmd == "version":
		cli.PrintVersion()
	default:
		cli.Usage()
	}
}

func drawRoutes(arg cli.Arg) {
	http.HandleFunc("/", root.Handle(arg))
	http.HandleFunc("/projects", projects.Handle(arg))
	if arg.MemberFile != "" {
		path := arg.Root + "/" + arg.MemberFile
		http.HandleFunc("/members", members.Handle(path))
	}

	data := template.Data{
		Name: arg.Name,
	}
	template.Init(rice.MustFindBox("embedded_assets/tmpl"), data)

	box := rice.MustFindBox("embedded_assets")
	embeddedFileServer := http.StripPrefix("/embedded_assets/", http.FileServer(box.HTTPBox()))
	http.Handle("/embedded_assets/", embeddedFileServer)
}
