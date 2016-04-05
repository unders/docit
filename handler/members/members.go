package members

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/unders/docit/template"
)

// Handle renders members parsed from the file.
// If parsing of member file fails, the prog will terminate.
func Handle(path string) func(w http.ResponseWriter, req *http.Request) {
	members := membersfromFile(path)

	return func(w http.ResponseWriter, req *http.Request) {
		template.RenderMembers(w, members)
	}
}

// membersFromFile will terminate the prog if it fails to
// parse the file.
func membersfromFile(path string) []template.Member {
	var members []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file %s, err %v", path, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Could not close file: %v, %v", file, err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		members = append(members, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Could not scan file %s, err %v", path, err)
	}

	length := len(members)

	tmplMembers := make([]template.Member, length)

	for i, mm := range members {
		tmplMembers[i].Email = mm
	}

	return tmplMembers
}
