#!/usr/bin/env bash

.PHONY:all
all: clean vendor fmt build

.PHONY:fmt
fmt:
	goimports -l -w -local "github.com/jiadas/" ./cmd ./internal ./pkg

.PHONY:vendor
vendor:
	govendor add +e
	govendor remove +u

.PHONY:clean
clean:
	rm -rf output/bin/

.PHONY:build
build:
	go build -o output/bin/helloweb ./cmd/helloweb


