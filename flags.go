package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Flags read from the command line.
var (
	index string
	root  string
	port  string
)

// Default flag values when not supplied from the command line.
const (
	indexDefault = "index.md"
	rootDefault  = "./"
	portDefault  = "80"
)

func usage() {
	fmt.Println("Usage: ")
	fmt.Println("")
	fmt.Printf("    %s serve -index=Readme.md -root=doc -port=8080\n", prog)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Printf("    %s serve \n", prog)
	fmt.Printf("    %s serve -index=Readme.md\n", prog)
	fmt.Printf("    %s serve -root=test\n", prog)
	fmt.Printf("    %s serve -index=index.md -root=doc\n", prog)
	fmt.Printf("    %s serve -index=index.md -root=doc -port=5000\n", prog)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Version")
	printVersion()
}

func printVersion() {
	fmt.Printf("    Version        : %s\n", Version)
	fmt.Printf("    UTC Build Time : %s\n", Buildstamp)
	fmt.Printf("    Git Commit Hash: %s\n", Githash)
}

func setFlags() *flag.FlagSet {
	f := flag.CommandLine
	f.Usage = usage

	f.StringVar(&index, "index", indexDefault, "Page to show for '/'")
	f.StringVar(&root, "root", rootDefault, "Root directory to serve files from")
	f.StringVar(&port, "port", portDefault, "The port")

	return f
}

func input() (string, context) {
	argLen := len(os.Args)

	f := setFlags()

	if argLen < 2 {
		return "", context{}
	}

	cmd := os.Args[1]
	if cmd != "serve" {
		return "", context{}
	}

	for _, flag := range os.Args[2:] {
		if !strings.HasPrefix(flag, "-") {
			fmt.Println("flag should be prefixed with a - sign: ", flag)
			return "", context{}
		}
	}

	_ = f.Parse(os.Args[2:])

	return cmd, context{index: index, root: root, port: port}
}
