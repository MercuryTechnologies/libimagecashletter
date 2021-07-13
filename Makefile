.PHONY: build

build:
	go fmt ./...
	@mkdir -p ./bin/
	CGO_ENABLED=1 go build -buildmode=c-shared -o ./bin/libimagecashletter.so ./
