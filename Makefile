build:
	@go build -o bin/lotofacil

run: build
	@./bin/lotofacil

test:
	@go test -v ./...