SHELL:=/bin/bash

.PHONY: install docs-to-proto run
install:
	go get -u github.com/NYTimes/openapi2proto/cmd/openapi2proto

docs-to-proto:
	openapi2proto -spec service.yaml -out service.proto

run: docs-to-proto
	./gen-protos.sh
