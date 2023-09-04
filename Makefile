.SUFFIXES:

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: t
t:
	go test ./...
