GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
CMD=cmd
BINARY_NAME=cmd

generate: build-linux build-mac

build:
		echo "Building Binary"
		$(GOBUILD) -o $(CMD) ./...

run:
		echo "Building and Executing Binary"
		$(GOBUILD) -o $(CMD) ./...
		$(CMD)/$(BINARY_NAME)

test:
		echo "Executing Tests"
		go test ./...