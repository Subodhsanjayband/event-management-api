build:
	@go build -o bin/event_manager

run:build
	@./bin/event_manager

test:
	@go test -v ./...
