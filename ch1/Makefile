.DEFAULT_GOAL := run

.PHONY:fmt vet build
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build -o hello
run: build
	./hello
clean: build
	go clean
