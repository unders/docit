.PHONY: help deps install gofmt oracle depgraph check release log shortlog clean
RELEASE=v1.0.4
VERSION="github.com/unders/docit/cli.Version=$(RELEASE)"
BUILDSTAMP="github.com/unders/docit/cli.Buildstamp=$(shell date -u '+%Y-%m-%dT%I:%M%p')"
GIT_HASH="github.com/unders/docit/cli.Githash=$(shell git rev-parse HEAD)"
LDFLAGS=-ldflags "-X $(VERSION) -X $(BUILDSTAMP) -X $(GIT_HASH)"
GOOS ?= darwin
GOARCH ?= amd64
PROG=out/docit_$(RELEASE)_$(GOOS)_$(GOARCH)

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
	godepgraph -s -horizontal github.com/unders/docit | dot -Tsvg -o doc/godepgraph.svg

check:
	gometalinter ./... --deadline=45s

release: clean check
	rice embed-go
	go build $(LDFLAGS) -o $(PROG)
	rm rice-box.go

log:
	@git log --graph --oneline --decorate

shortlog:
	@git shortlog

clean:
	@if [ -f $(PROG) ] ; then rm $(PROG) ; fi
