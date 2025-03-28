#!/usr/bin/make -f

test: fmt
	GORACE="atexit_sleep_ms=50" go test -count=1 -timeout=5s -short -race -covermode=atomic ./...

fmt:
	go mod tidy && go fmt ./...

compile:
	go build ./...

build: test compile

.PHONY: test fmt compile build
