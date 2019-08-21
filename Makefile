.PHONY: all clean build install test tf cover protofmt protoc protolint protodocs import-spec

# make sure we turn on go modules
export GO111MODULE := on

TOOLS := cmd/customd cmd/customcli

# MODE=count records heat map in test coverage
# MODE=set just records which lines were hit by one test
MODE ?= set

# Check if linter exists
LINT := $(shell command -v golangci-lint 2> /dev/null)

# for dockerized prototool
USER := $(shell id -u):$(shell id -g)
DOCKER_BASE := docker run --rm --user=${USER} -v $(shell pwd):/work iov1/prototool:v0.2.2
PROTOTOOL := $(DOCKER_BASE) prototool
PROTOC := $(DOCKER_BASE) protoc
WEAVEDIR=$(shell go list -m -f '{{.Dir}}' github.com/iov-one/weave)

all: import-spec test lint install

build:
	go build $(BUILD_FLAGS) -o $(BUILDOUT) ./cmd/customd

dist:
	cd cmd/bnsd && $(MAKE) dist

install:
	for ex in $(TOOLS); do cd $$ex && make install && cd -; done

test:
	@# customd binary is required by some tests. In order to not skip them, ensure customd binary is provided and in the latest version.
	go vet -mod=readonly ./...
	go test -mod=readonly -race ./...

lint:
	@go mod vendor
	docker run --rm -it -v $(shell pwd):/go/src/github.com/iov-one/weave-starter-kit -w="/go/src/github.com/iov-one/weave-starter-kit" golangci/golangci-lint:v1.17.1 golangci-lint run ./...
	@rm -rf vendor

# Test fast
tf:
	go test -short ./...

test-verbose:
	go vet ./...
	go test -v -race ./...

mod:
	go mod tidy

# TODO write github.com/iov-one/weave-starter-kit/cmd/customd/client and scenarios, here when implemented \
	go test -mod=readonly -covermode=$(MODE) \
		-coverpkg=github.com/iov-one/weave/cmd/customd/app,\
		-coverprofile=coverage/custonmd_scenarios.out \
		github.com/iov-one/weave-starter-kit/cmd/bnsd/scenarios
# TODO \
	go test -mod=readonly -covermode=$(MODE) \
		-coverpkg=github.com/iov-one/weave-starter-kit/cmd/bnsd/app,github.com/iov-one/weave-starter-kit/cmd/bnsd/client,github.com/iov-one/weave-starter-kit/app \
		-coverprofile=coverage/bnsd_client.out \
		github.com/iov-one/weave-starter-kit/cmd/bnsd/client
cover:
	@# TODO write github.com/iov-one/weave-starter-kit/cmd/bnsd/client when implemented
	@go test -mod=readonly -covermode=$(MODE) \
		-coverpkg=github.com/iov-one/weave-starter-kit/cmd/customd/app, \
		-coverprofile=coverage/customd_app.out \
		github.com/iov-one/weave-starter-kit/cmd/customd/app
		cat coverage/*.out > coverage/coverage.txt

novendor:
	@rm -rf ./vendor

protolint: novendor
	$(PROTOTOOL) lint

protofmt: novendor
	$(PROTOTOOL) format -w

protodocs: novendor
	./scripts/build_protodocs.sh docs/proto

protoc: protolint protofmt protodocs
	$(PROTOTOOL) generate

import-spec:
	@rm -rf ./spec
	@mkdir -p spec/github.com/iov-one/weave
	@cp -r ${WEAVEDIR}/spec/gogo/* spec/github.com/iov-one/weave
	@chmod -R +w spec

inittm:
	tendermint init --home ~/.custom

runtm:
	tendermint node --home ~/.custom > ~/.custom/tendermint.log &
