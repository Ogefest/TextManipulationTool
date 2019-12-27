# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=bin/fmt
BINARY_UNIX=$(BINARY_NAME)_unix
    
all: build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v src/*.go

#test: 
#	$(GOTEST) -v ./...
#

install:
	cp $(BINARY_NAME) /usr/bin/fmt

uninstall:
	rm -f /usr/bin/fmt

