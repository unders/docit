package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/unders/docit/cli"
)

func serve(arg cli.Arg) {
	fmt.Printf("Serving static files on 0.0.0.0:%s ...\n", arg.Port)
	log.Fatal(http.ListenAndServe(":"+arg.Port, http.FileServer(http.Dir(arg.Root))))
}
