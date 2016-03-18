package main

var (
	// Version the version of the prog.
	Version = "No Version Provided"
	// Buildstamp the time the prog was built.
	Buildstamp = "No Buildstamp provided"
	// Githash the git commit hash
	Githash = "No Githash provided"
	prog    = "docit"
)

type context struct {
	index string
	root  string
	port  string
}

func main() {
	switch command, ctx := input(); true {
	case command == "serve":
		serve(ctx)
	default:
		usage()
	}
}
