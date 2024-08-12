build:
	go build

run: build
	./http_server

test:
	go test -v ./...