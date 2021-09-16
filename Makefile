GOPATH ?= $(HOME)/go
SRCPATH := $(patsubst %/,%,$(GOPATH))/src

default: install

.PHONY: options
options:
	protoc -I. -I$(SRCPATH) -I./vendor  \
		--go_out=:$(SRCPATH) \
		options/query_validate.proto

.PHONY: install
install:
	go install

.PHONY: example
example: default
	protoc -I. -I${SRCPATH} -I./vendor -I./vendor/github.com/grpc-ecosystem/grpc-gateway/v2 --atlas-query-validate_out=$(GOPATH)/src/ example/example.proto

test: example
	go test  ./...

.PHONY: vendor
vendor:
	dep ensure -vendor-only

.PHONY: gentool
gentool:
	docker build -t infoblox/atlas-gentool:atlas-validate-query-dev .
	docker image prune -f --filter label=stage=server-intermediate
