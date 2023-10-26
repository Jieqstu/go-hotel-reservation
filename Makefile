 build:
	@go build -o bin/api

run: build
	@./bin/api --listenAddr :5010

seed:
	@go run scripts/seed.go

test:
	@go test -v ./...