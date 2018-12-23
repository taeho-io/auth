# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: proto test

proto:
	@go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . auth.proto --go_out=plugins=grpc:${GOPATH}/src

build: proto
	go build -o build/auth cmd/main.go
    
test:
	@go get github.com/rakyll/gotest
	gotest -p 1 -race -cover -v ./...

lint:
	golangci-lint run ./...

.PHONY: generate_mocks
generate_mocks:
	@go get github.com/vektra/mockery
	mockery -dir=./pkg/token -name=Token

.PHONY: clean_mocks
clean_mocks:
	find . -name "mock_*.go" -type f -delete
	rm -rf mocks
