VERSION := $(shell git describe --tag --always --dirty)

.PHONY: setup
setup: ## Install all the build and lint dependencies
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.15.0
	GO111MODULE=off go get -u golang.org/x/tools/cmd/cover

.PHONY: test
test: ## Run all the tests
	echo 'mode: atomic' > coverage.txt && GO111MODULE=on go test -covermode=atomic -coverpkg=./... -coverprofile=coverage.txt -race -timeout=30s ./...

.PHONY: fasttest
fasttest: ## Run short tests
	echo 'mode: atomic' > coverage.txt && GO111MODULE=on go test -short -covermode=atomic -coverprofile=coverage.txt -race -timeout=30s ./...

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt


.PHONY: lint
lint: ## Run all the linters
	golangci-lint run 

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: build
build: ## Build a version
	GO111MODULE=on  go build -tags=jsoniter -v ./cmd/...

.PHONY: clean
clean: ## Remove temporary files
	go clean

.PHONY: install
install: ## install project and it's dependancies, useful for autocompletion feature
	go install -i


.PHONY: version
version: ## display version of augeas
	@echo $(VERSION)

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
