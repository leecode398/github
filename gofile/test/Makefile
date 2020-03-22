GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=main

all: run
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GORUN) sock.go
sock:
	$(GORUN) sock.go
client:
	$(GORUN) client.go
sql:
	$(GORUN) mysql.go
ref:
	$(GORUN) reflect.go


