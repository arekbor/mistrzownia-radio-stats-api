BIN_NAME = mistrzownia-converter-stats-api

build:
	@go build -o ./bin/$(BIN_NAME) ./cmd/main.go

run: build
	@./bin/$(BIN_NAME)