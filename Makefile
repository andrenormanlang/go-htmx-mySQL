# Go parameters
GOCMD=go
TEMPL=templ
BUILD_DIR=./tmp
CMSGO_DIR=./cmd/cmsgo
CMSGO_ADMIN_DIR=./cmd/cmsgo-admin

# Name of the binary
BINARY_NAME=cmsgo.exe
ADMIN_BINARY_NAME=cmsgo-admin

all: build test

build:
	$(TEMPL) generate
	GIN_MODE=release $(GOCMD) build -ldflags "-s" -v -o $(BUILD_DIR)/$(BINARY_NAME) $(CMSGO_DIR)
	GIN_MODE=release $(GOCMD) build -ldflags "-s" -v -o $(BUILD_DIR)/$(ADMIN_BINARY_NAME) $(CMSGO_ADMIN_DIR)

test:
	$(GOCMD) test -v ./...

clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)

install-tools:
	go install github.com/a-h/templ/cmd/templ@v0.2.543

.PHONY: all build test clean