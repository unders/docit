.PHONY: help install release clean
VERSION="github.com/unders/docit/cli.Version=v1.0.0"
BUILDSTAMP="github.com/unders/docit/cli.Buildstamp=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')"
GIT_HASH="github.com/unders/docit/cli.Githash=$(shell git rev-parse HEAD)"
LDFLAGS=-ldflags "-X $(VERSION) -X $(BUILDSTAMP) -X $(GIT_HASH)"
PROG=out/docit

help:
	@echo "Usage:"
	@echo "        make command"
	@echo ""
	@cat Makefile

install:
	@./deps install

release: clean
	go build $(LDFLAGS) -o $(PROG)

clean:
	@if [ -f $(PROG) ] ; then rm $(PROG) ; fi
