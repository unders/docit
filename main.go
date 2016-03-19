package main

import "github.com/unders/docit/cli"

func main() {
	switch cmd, arg := cli.Parse(); true {
	case cmd == "serve":
		serve(arg)
	case cmd == "version":
		cli.PrintVersion()
	default:
		cli.Usage()
	}
}
