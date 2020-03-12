GOROOT := /usr/local/go
GOPATH := ${HOME}/go

build:
	@echo "Building main_go..."
	GOROOT=$(GOROOT) GOPATH=$(GOPATH) go build -o main_go ./main.go

run: build
	@echo "Running main_go..."
	./main_go

.DEFAULT_GOAL := run