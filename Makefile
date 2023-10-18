build:
	@go build -o bin/api

run: build
	@./bin/api --listenAddr :5010

test:
	@go test -v ./...