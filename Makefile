BINARY      := data-bin
BUILD_DIR   := build
ZIP         := $(BUILD_DIR)/$(BINARY).zip
LAMBDA_MAIN := ./cmd/$(BINARY)
TF_DIR      := deployment

GOARCH      ?= arm64
GOOS        ?= linux

.PHONY: all build zip test lint clean deploy destroy tf-init tf-plan

all: build zip

## build: compile the Lambda binary
build:
	@mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -tags lambda.norpc -ldflags="-s -w -linkmode=internal" -o $(BUILD_DIR)/bootstrap $(LAMBDA_MAIN)

## zip: package the binary for Lambda deployment
zip: build
	cd $(BUILD_DIR) && zip -j $(BINARY).zip bootstrap

## test: run all tests
test:
	go test ./...

## lint: run golangci-lint
lint:
	golangci-lint run ./...

## clean: remove build artifacts
clean:
	rm -rf $(BUILD_DIR)

## tf-init: initialise Terraform (run once)
tf-init:
	terraform -chdir=$(TF_DIR) init

## tf-plan: show Terraform plan
tf-plan: zip
	terraform -chdir=$(TF_DIR) plan

## deploy: build, zip, then apply Terraform
deploy: zip
	terraform -chdir=$(TF_DIR) apply -auto-approve

## destroy: tear down all AWS resources
destroy:
	terraform -chdir=$(TF_DIR) destroy

## help: list available targets
help:
	@grep -E '^##' Makefile | sed 's/## //'
