.PHONY: lint pretty check-pretty

IMAGE		= satu-dental/be
DEPLOY_DIRS	= $(shell cd deploy && ls -d * | grep -v "_output")
export CI_COMMIT_SHA ?= $(shell git log --format="%H" -n 1)
export LINTER_VERSION ?= 1.56.2
export GOSEC_VERSION ?= 2.18.2
VERSION 	= ${CI_COMMIT_SHA}
OUTPUT_DIR := deploy/_output

dep:
	go mod download
	go mod tidy

wire:
	wire ./...
