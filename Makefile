include .env

GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOINSTALL = $(GOCMD) install
GOGET = $(GOCMD) get
BINARY_NAME = pandp


all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) main.go

test:
	$(GOTEST) ./...

install:
	$(GOINSTALL)

clean:
	$(GOCMD) clean
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) main.go
	./$(BINARY_NAME)