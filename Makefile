BINARY_NAME=gotime
INSTALL_DIR=/usr/local/bin

KERNEL := $(shell uname | awk '{print tolower($0)}')

# BUILD_ARGS=-linkshared


ifeq ($(shell uname -m),arm64)
    ARCH := arm64
else
    ARCH := amd64
endif

tests:
	go test ./...

build:
	go build ${BUILD_ARGS} -o dist/${BINARY_NAME}-${KERNEL}-${ARCH} main.go

build-all:
	GOARCH=amd64 GOOS=darwin go build ${BUILD_ARGS} -o dist/${BINARY_NAME}-darwin-amd64 main.go
	GOARCH=amd64 GOOS=linux go build ${BUILD_ARGS} -o dist/${BINARY_NAME}-linux-amd64 main.go

	GOARCH=amd64 GOOS=darwin go build ${BUILD_ARGS} -o dist/${BINARY_NAME}-darwin-amd64 main.go
	GOARCH=arm64 GOOS=linux go build ${BUILD_ARGS} -o dist/${BINARY_NAME}-linux-arm64 main.go

clean:
	go clean
	rm dist/*

install: build
	cp dist/${BINARY_NAME}-${KERNEL}-${ARCH} ${INSTALL_DIR}/${BINARY_NAME}

uninstall:
	rm ${INSTALL_DIR}/${BINARY_NAME}