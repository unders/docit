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
		template.Init(rice.MustFindBox("embedded_assets/tmpl"), arg.Name)

		routes(arg)

		fmt.Printf("Serve files at http://0.0.0.0:%s from %s\n", arg.Port, arg.Root)
		log.Fatal(http.ListenAndServe(":"+arg.Port, nil))

	case cmd == "version":
		cli.PrintVersion()
	default:
		cli.Usage()
	}
}

func routes(arg cli.Arg) {
	http.HandleFunc("/", root.Handle(arg))
	http.HandleFunc("/projects", projects.Handle(arg.Root))
	if arg.MemberFile != "" {
		path := arg.Root + "/" + arg.MemberFile
		http.HandleFunc("/members", members.Handle(path))
	}

	fileServer := http.FileServer(rice.MustFindBox("embedded_assets").HTTPBox())
	h := http.StripPrefix("/embedded_assets/", fileServer)
	http.Handle("/embedded_assets/", h)
}
