.PHONY: help install release
VERSION=v1.0.0
TIME=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
GIT_HASH=$(shell git rev-parse HEAD)

help:
	@echo "Usage:"
	@echo "        make command"
	@echo ""
	@cat Makefile

install:
	@./deps install

release:
	go build -o out/docit -ldflags "-X main.Version=$(VERSION) -X main.Buildstamp=$(TIME) -X main.Githash=$(GIT_HASH)"
