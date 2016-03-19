.PHONY: help deps install gofmt oracle depgraph check release clean
VERSION="github.com/unders/docit/cli.Version=v1.0.0"
BUILDSTAMP="github.com/unders/docit/cli.Buildstamp=$(shell date -u '+%Y-%m-%dT%I:%M%p')"
GIT_HASH="github.com/unders/docit/cli.Githash=$(shell git rev-parse HEAD)"
LDFLAGS=-ldflags "-X $(VERSION) -X $(BUILDSTAMP) -X $(GIT_HASH)"
PROG=out/docit

help:
	@cat Makefile

deps:
	@bin/deps install

install:
	go install $(LDFLAGS)

gofmt:
	gofmt -l -s -w .

oracle:
	pythia github.com/unders/docit

depgraph:
	godepgraph -horizontal github.com/unders/docit | dot -Tsvg -o doc/godepgraph.svg

check:
	@ gofmt -l . | grep -vE ''
	gometalinter ./... --deadline=10s

release: clean check
	go build $(LDFLAGS) -o $(PROG)

clean:
	@if [ -f $(PROG) ] ; then rm $(PROG) ; fi
