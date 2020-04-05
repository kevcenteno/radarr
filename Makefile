# Project variables
NAME        := radarr
VENDOR      := SkYNewZ
DESCRIPTION := Radarr Go client
MAINTAINER  := Quentin Lemaire <quentin@lemairepro.fr>
URL         := https://github.com/$(VENDOR)/$(NAME)
LICENSE     := GPL-3
PACKAGE			:= github.com/$(VENDOR)/$(NAME)
BINARY_DIR	:= cmd/radarr

# Build variables
BUILD_DIR   := bin
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
VERSION     ?= $(shell git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.0-$(COMMIT_HASH)")

# Go variables
GOCMD       := GO111MODULE=on go
GOOS        ?= $(shell go env GOOS)
GOARCH      ?= $(shell go env GOARCH)
GOPKGS      ?= $(shell $(GOCMD) list ./... | grep -v /vendor)

GOBUILD     ?= CGO_ENABLED=0 $(GOCMD) build

.PHONY: all
all: clean test build

#########################
## Development targets ##
#########################
.PHONY: clean
clean: ## Clean workspace
	@ $(MAKE) --no-print-directory log-$@
	rm -rf ./$(BUILD_DIR)
	rm -rf ./$(NAME)

.PHONY: test
test: ## Run tests
	@ $(MAKE) --no-print-directory log-$@
	$(GOCMD) test -v $(GOPKGS) -run "Test" -coverprofile=cover.out -cover

.PHONY: verify
verify: ## Verify 'vendor' dependencies
	@ $(MAKE) --no-print-directory log-$@
	$(GOCMD) mod verify

###################
## Build targets ##
###################
.PHONY: build
build: clean ## Build binary for current OS/ARCH
	@ $(MAKE) --no-print-directory log-$@
	cd $(BINARY_DIR) && \
	$(GOBUILD) -ldflags "-X main.Version=$(VERSION)" -o ../../$(BUILD_DIR)/$(NAME)

.PHONY: build-all
build-all: GOOS      = linux darwin
build-all: GOARCH    = amd64
build-all: clean gox ## Build binary for all OS/ARCH
	@ $(MAKE) --no-print-directory log-$@
	cd $(BINARY_DIR) && \
	gox -arch="$(GOARCH)" -os="$(GOOS)" -output="../../$(BUILD_DIR)/{{.Dir}}-$(VERSION)-{{.OS}}-{{.Arch}}" -ldflags "-X main.Version=$(VERSION)"

####################
## Helper targets ##
####################
.PHONY: gox
gox: ## Installing gox for cross compile
	@ $(MAKE) --no-print-directory log-$@
	GO111MODULE=off go get -u github.com/mitchellh/gox

########################################################################
## Self-Documenting Makefile Help                                     ##
## https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html ##
########################################################################
.PHONY: help
help:
	@ grep -h -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

log-%:
	@ grep -h -E '^$*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m==> %s\033[0m\n", $$2}'
