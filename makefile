.PHONY: build clean default

build: clean
	@go build -o bin/machineid ./cli/machineid/main.go

clean:
	@rm -rf ./bin/*

default: build
