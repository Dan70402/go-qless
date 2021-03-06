PROJ=$(shell realpath $$PWD/../../../..)
ENV=env GOPATH=$(PROJ)
CLI=$(ENV) easyjson -all

JSON_SRC_FILES=\
	structs.go

JSON_OUT_FILES:=$(JSON_SRC_FILES:%.go=%_easyjson.go)

.PHONY: fmt

json: $(JSON_OUT_FILES)

%_easyjson.go: %.go
	$(CLI) $^

fmt:
	$(ENV) goimports -w $$(find . -path ./vendor -prune -o -depth 1 -name "*.go" -print)
