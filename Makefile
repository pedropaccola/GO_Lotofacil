build:
	@go build -o bin/lotofacil
	@GOOS=windows go build -o bin/lotofacil-win.exe

run: build
	@./bin/lotofacil

test:
	@go test -v ./...