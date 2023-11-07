build:
	@go build -o bin/maingo

run: build
	@./bin/maingo

test:
	@go test -v ./...
