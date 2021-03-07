SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=

export PATH := ./bin:$(PATH)
export GO111MODULE := on
export GOPROXY = https://proxy.golang.org,direct

build:
	go build
.PHONY: build

.DEFAULT_GOAL := build
