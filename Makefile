GOPATH ?= $(HOME)/go
SRCPATH := $(patsubst %/,%,$(GOPATH))/src

default: options install

.PHONY: options
options:
	protoc -I. -I$(SRCPATH) -I./vendor  \
		--gogo_out="Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:$(SRCPATH)" \
		options/query_validate.proto

.PHONY: install
install:
	go install

.PHONY: example
example: default
	protoc -I. -I${SRCPATH} -I./vendor -I./vendor/github.com/grpc-ecosystem/grpc-gateway --atlas-query-validate_out=. example/example.proto

test: example
	go test  ./...
	
.PHONY: vendor
vendor:
	dep ensure -vendor-only