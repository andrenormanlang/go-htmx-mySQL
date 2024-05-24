GOCMD=go
TEMPL=templ
BUILD_DIR=./tmp
CMSGO_DIR=./cmd/cmsgo
CMSGO_ADMIN_DIR=./cmd/cmsgo-admin

# Name of the binary
BINARY_NAME=cmsgo
ADMIN_BINARY_NAME=cmsgo-admin

all: build test

build:
	$(TEMPL) generate
	$(GOCMD) build -v -o $(BUILD_DIR)/$(BINARY_NAME) $(CMSGO_DIR)
	$(GOCMD) build -v -o $(BUILD_DIR)/$(ADMIN_BINARY_NAME) $(CMSGO_ADMIN_DIR)

test:
	$(GOCMD) test -v ./...

clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)

.PHONY: all build test clean