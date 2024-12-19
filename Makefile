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

prepare_env:
	cp -r migrations tests/helpers/

build: prepare_env install-tailwindcss
	$(TEMPL) generate
	fuser -k 8080/tcp || true # Kill process using port 8080
	GIN_MODE=release $(GOCMD) build -ldflags "-s" -v -o $(BUILD_DIR)/$(BINARY_NAME) $(CMSGO_DIR)
	GIN_MODE=release $(GOCMD) build -ldflags "-s" -v -o $(BUILD_DIR)/$(ADMIN_BINARY_NAME) $(CMSGO_ADMIN_DIR)
	./tailwindcss -i ./static/css/custom.css -o ./static/css/style.css --minify


test:prepare_env
	$(GOCMD) test -v ./...


clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)

install-tools:
	go install github.com/pressly/goose/v3/cmd/goose@v3.18.0
	go install github.com/a-h/templ/cmd/templ@v0.2.543 
	go install github.com/cosmtrek/air@v1.49.0

install-tailwindcss:
	if [ ! -f tailwindcss ]; then \
		curl -sL https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.16/tailwindcss-linux-x64 -o tailwindcss-linux-x64 \
		&& echo "33f254b54c8754f16efbe2be1de38ca25192630dc36f164595a770d4bbf4d893  tailwindcss-linux-x64" | sha256sum -c \
		&& chmod +x tailwindcss-linux-x64 \
		&& mv tailwindcss-linux-x64 tailwindcss; \
	fi

.PHONY: all build test clean