SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
.DEFAULT_GOAL := help
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

GO111MODULES=on
GIT_TAG := $(shell git describe --always --abbrev=0 --tags)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT := $(shell git log --pretty=format:'%h' -n 1)
VERSION="$(GIT_TAG)-$(GIT_BRANCH).$(GIT_COMMIT)"

ASCIIFY := "asciify.bin"
ASCIIFY_PKG_BUILD := "./cmd/asciify"
RELEASE_ZIP := "asciify.zip"

.PHONY: all build release
all: build release
build: $(ASCIIFY) ## Build Binary
release: $(RELEASE_ZIP) ## Package release artifact

$(ASCIIFY): dep
	@echo "🍳 Building $(ASCIIFY)"
	@go build -i -v -o $(ASCIIFY) -ldflags "-X main.version=$(GIT_TAG)-$(GIT_BRANCH).$(GIT_COMMIT)" $(ASCIIFY_PKG_BUILD)

$(RELEASE_ZIP): $(ASCIIFY)
	@echo "🍳 Building $(RELEASE_ZIP)"
	zip --junk-paths $(RELEASE_ZIP) $(ASCIIFY) README.md

.PHONY:clean
clean: ## Remove previous builds
	@echo "🧹 Cleaning old build"
	@go clean
	@rm -f $(ASCIIFY) $(RELEASE_ZIP)

.PHONY: dep
dep: ## go get all dependencies
	@echo "🛎 Updatig Dependencies"
	@go get -v -d ./...

.PHONY: run
run: dep ## Compiles and runs Binary
	@go run -race $(ASCIIFY_PKG_BUILD) --config configs/config.toml

.PHONY: help
help:  ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## Runs go test with default values
	@echo "🍜 Testing $(ASCIIFY)" @go test -v -count=1 -race ./...
