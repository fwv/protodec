.PHONY: all build 
BIN_FILE=protodesc
all: build
build:
	@go build -o "bin/${BIN_FILE}" cmd/main.go
