lint:
	@echo "Checking if golangci-lint is installed..."
	@if ! command -v golangci-lint > /dev/null; then \
		echo "golangci-lint not found. Installing..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.0.2; \
	fi
	@echo "Running linter..."
	@golangci-lint run ./... --fix --verbose --config .golangci.yaml

align:
	@echo "Checking if betteralign is installed..."
	@if ! command -v betteralign > /dev/null; then \
		echo "betteralign not found. Installing..."; \
		go install github.com/dkorunic/betteralign/cmd/betteralign@latest; \
	fi
	@echo "Running align..."

	@betteralign -apply ./...

.PHONY: lint align
