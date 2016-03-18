package main

import (
	"fmt"
	"log"
	"net/http"
)

func serve(ctx context) {
	fmt.Printf("ctx: %#v\n", ctx)

	fmt.Printf("Serving static files on 0.0.0.0:%s ...\n", ctx.port)
	log.Fatal(http.ListenAndServe(":"+ctx.port, http.FileServer(http.Dir(ctx.root))))
}
