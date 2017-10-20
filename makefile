.PHONY: build clean default test

build: clean
	@go build -o bin/machineid ./cli/machineid/main.go

clean:
	@rm -rf ./bin/*

test:
	go test ./...

default: build
