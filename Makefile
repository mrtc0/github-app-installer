VERSION := $(shell git describe --tags)
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION)"
GO111MODULE=on

GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin
GOSRC=$(GOPATH)/src

.PHONY: build
build:
	go build $(LDFLAGS) ./cmd/github-app-installer
