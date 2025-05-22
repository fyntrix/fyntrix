unit_test:
	@echo "Running unit tests..."
	@go test ./...

race_test:
	@echo "Running race condition tests..."
	@go test ./... -race

fmt:
	@echo "Formatting code..."
	@go tool gofumpt -l -w .

lint:
	@echo "Running lint..."
	@go tool golangci-lint  run ./... --timeout=20m0s

check: fmt lint

.PHONY: test unit_test race_test
.PHONY: fmt lint check