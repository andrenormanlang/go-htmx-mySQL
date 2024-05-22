GOCMD=go
TEMPL=templ 
BUILD_DIR=./tmp
CMSGO_DIR=./cmd/cmsgo

BINARY_NAME=cmsgo

all: build test

build:
	$(GOCMD) build -v -o $(BUILD_DIR)/$(BINARY_NAME) $(CMSGO_DIR)

test:
	$(TEMPL) generate
	$(GOCMD) test -v ./...

clean:
	$(GOCMD) clean
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

run:
	$(GOCMD) run $(CMSGO_DIR)/main.go
    

.PHONY: all build test clean run



