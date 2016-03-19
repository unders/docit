package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/unders/docit/cli"
)

func serve(arg cli.Arg) {
	fmt.Printf("Serving static files at http://0.0.0.0:%s from dir %s\n", arg.Port, arg.Root)
	log.Fatal(http.ListenAndServe(":"+arg.Port, http.FileServer(http.Dir(arg.Root))))
}
