NAME = vertex
VERSION = none

.PHONY: example release

example:
	go run example/example.go

release:
	GOPROXY=proxy.golang.org go list -m github.com/dimkauzh/vertex@$(VERSION)
