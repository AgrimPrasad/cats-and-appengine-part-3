#!/bin/bash

: ${GOPATH:?"ERROR: GOPATH not set"}

export PATH=$PATH:$GOPATH/bin

protoc_path=$(which protoc)
if [ ! -x "$protoc_path" ]; then
	echo "ERROR: Protobuf Compiler not found"
	exit 1
fi

protoc_gen_go_path=$(which protoc-gen-go)
if [ ! -x "$protoc_gen_go_path" ]; then
	echo "ERROR: Protobuf Go bindings not found"
	echo "go get -u github.com/golang/protobuf/{proto,protoc-gen-go}"
	exit 1
fi

protoc --go_out=. *.proto
