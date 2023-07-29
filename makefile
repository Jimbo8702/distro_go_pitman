build: 
	@go build -o bin/gocrawl

run: build
	@./build/gocrawl

test:
	@go test -v ./...


