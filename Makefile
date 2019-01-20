
PROJECT_ROOT := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
BUILD_ROOT := $(PROJECT_ROOT)/bin

all: build test-cover

.PHONY: build
build: build-linux build-mac

.PHONY: pack
pack: pack-linux pack-mac

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_ROOT)/linux-amd64/cdl -x ./

.PHONY: pack-linux
pack-linux: build-linux
	zip -j $(BUILD_ROOT)/cdl-linux-amd64.zip $(BUILD_ROOT)/linux-amd64/cdl

.PHONY: build-mac
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_ROOT)/darwin-amd64/cdl -x ./

.PHONY: pack-mac
pack-mac: build-mac
	zip -j $(BUILD_ROOT)/cdl-darwin-amd64.zip $(BUILD_ROOT)/darwin-amd64/cdl

.PHONY: test
test:
	go test -covermode=count -coverprofile=coverage.out ./...

.PHONY: test-cover
test-cover: test
	go tool cover -html=coverage.out

