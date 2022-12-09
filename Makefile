HAS_GINKGO := $(shell command -v ginkgo;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_COUNTERFEITER := $(shell command -v counterfeiter;)
PLATFORM := $(shell uname -s)

# #### DEPS ####
.PHONY: deps-counterfeiter deps-ginkgo deps-modules

deps-counterfeiter:
ifndef HAS_COUNTERFEITER
	go install github.com/maxbrunsfeld/counterfeiter/v6@latest
endif

deps-ginkgo:
ifndef HAS_GINKGO
	go install github.com/onsi/ginkgo/v2/ginkgo
endif

deps-modules:
	go mod download

# #### TEST ####
.PHONY: lint test

lint:
ifndef HAS_GOLANGCI_LINT
ifeq ($(PLATFORM), Darwin)
	brew install golangci-lint
endif
ifeq ($(PLATFORM), Linux)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
endif
endif
	golangci-lint run

API_KEY := $(or $(API_KEY), $(shell op read "op://Private/NASA API/credential"))
test/inputs/config.yaml: test/inputs/config-template.yaml
	ytt --file test/inputs/config-template.yaml \
		--data-value apiKey=${API_KEY} \
		> test/inputs/config.yaml

test: deps-modules deps-ginkgo test/inputs/config.yaml
	ginkgo -r .

# #### BUILD ####
.PHONY: build
SOURCES = $(shell find . -name "*.go" | grep -v "_test\." )
VERSION := $(or $(VERSION), dev)
LDFLAGS="-X github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day/cmd.Version=$(VERSION)"

build: build/nasa-image-of-the-day

build/nasa-image-of-the-day: $(SOURCES) deps-modules
	go build -o $@ -ldflags ${LDFLAGS} github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day

build-all: build/nasa-image-of-the-day-arm6 build/nasa-image-of-the-day-arm7 build/nasa-image-of-the-day-darwin-amd64

build/nasa-image-of-the-day-arm6: $(SOURCES) deps-modules
	GOOS=linux GOARCH=arm GOARM=6 go build -o $@ -ldflags ${LDFLAGS} github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day

build/nasa-image-of-the-day-arm7: $(SOURCES) deps-modules
	GOOS=linux GOARCH=arm GOARM=7 go build -o $@ -ldflags ${LDFLAGS} github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day

build/nasa-image-of-the-day-darwin-amd64: $(SOURCES) deps-modules
	GOOS=darwin GOARCH=amd64 go build -o $@ -ldflags ${LDFLAGS} github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day
