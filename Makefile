build:
	@go build -o bin/go-auth-api cmd/main.go

test:
	@go test -v ./...

run:
	@./bin/go-auth-api