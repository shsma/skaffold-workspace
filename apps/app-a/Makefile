

SHELL := /bin/bash

OS ?= $(shell uname | tr '[:upper:]' '[:lower:]')
ARCH ?= $(shell uname -m)
VERSION ?= $(shell git rev-parse --short HEAD)
BUILD_DATE?=$(shell date -u +%FT%TZ)
COMMIT ?= $(shell git rev-parse --short HEAD)

LDFLAGS := -ldflags '-X main.Version=$(VERSION)-$(OS)-$(ARCH) -X main.BuildDate=$(BUILD_DATE) -X main.Commit=$(COMMIT)'

.PHONY: install-deps
install-deps:
	go mod download -x

api-server: $(API_SERVER_REQUIRES)
	cd cmd/ && CGO_ENABLED=0 go build $(LDFLAGS) -o ../$@ main.go

api-server-amd64: $(API_SERVER_REQUIRES)
	cd cmd/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o ../$@ main.go
