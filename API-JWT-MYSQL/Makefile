.PHONY: build test run clean

build:
	go build -o bin/ecom cmd/main.go

test:
	go test -v ./..

run: build
	./bin/ecom

clean:
	rm -rf bin/ecom
