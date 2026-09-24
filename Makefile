.DEFAULT_GOAL := all

BUILD_DIR ?= ./bin

healthcheck:
	curl -i localhost:4000/v1/healthcheck

fmt:
	go fmt ./...

lint: fmt
	golangci-lint run

vet: lint
	go vet ./...

test:
	go test -v -cover ./...

bench:
	go test -bench . ./...

clean:
	rm -r $(BUILD_DIR)

api: vet
	go build -o $(BUILD_DIR)/api ./cmd/api

all: api

.PHONY: all api healthcheck vet lint fmt test bench clean
