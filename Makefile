.PHONY: all one two

GOPATH := $(or $(GOPATH), $(HOME)/go)
DIST_DIR := out
WIRE := $(GOPATH)/bin/wire

$(WIRE):
	GOPATH=$(GOPATH) go install -mod=mod github.com/google/wire/cmd/wire

default: di build;

di: $(WIRE)
	$(WIRE) gen -tags dynamic -output_file_prefix build_server_  ./cmd

build: 
	go build -tags dynamic -o $(DIST_DIR) ./cmd

test-unit:
	go test -tags dynamic -parallel=1 -count=1 -coverprofile=./c.out -v ./...

test-cover:
	go tool cover -func ./c.out

test-cover-html:
	go tool cover -html=c.out -o cover.html

run:
	./out -f kafka.conf

up-depend:
	go get -u github.com/Tracking-SYS/proto-tracking-gen
	go get -u github.com/Tracking-SYS/go-lib
	go mod tidy
	go mod vendor