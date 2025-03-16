.DEFAULT_GOAL := build

.PHONY: format
format:
	@echo "Formatting the project..."
	@go fmt ./...

.PHONY: vet
vet: format
	@echo "Vetting the project..."
	@go vet ./...

.PHONY: build
build: vet
	@echo "Building the project..."
	@go build

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm go-basics-makefile
