GO ?= go

.PHONY: debug
debug:
	 $(GO) build -o ${GOPATH}/bin/i18n-stringer i18n-stringer.go
	 $(GO) generate ./...