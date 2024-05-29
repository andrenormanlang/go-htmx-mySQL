# Version: 0.3

# We base our image from the Ubuntu light image
FROM ubuntu:jammy-20240125

# Environment variables needed for the build system
ENV TZ=Europe/Stockholm
ENV DEBIAN_FRONTEND="noninteractive"
ENV PATH="${PATH}:/usr/local/go/bin"

# Identify the maintainer of an image
LABEL maintainer="andrenormanlang@gmail.com"

# Install build dependencies
RUN apt-get update \
    && apt-get install -y --no-install-recommends --no-install-suggests \
    ca-certificates \
    git \
    make \
    pkg-config \
    wget

# Install Go
RUN wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz -P /tmp \
    && tar -C /usr/local -xzf /tmp/go1.22.0.linux-amd64.tar.gz \
    && rm /tmp/go1.22.0.linux-amd64.tar.gz

# Install Go tools
RUN GOBIN=/usr/local/bin go install github.com/pressly/goose/v3/cmd/goose@v3.18.0 \
    && GOBIN=/usr/local/bin go install github.com/a-h/templ/cmd/templ@v0.2.543 \
    && GOBIN=/usr/local/bin go install github.com/cosmtrek/air@v1.49.0

# Install golangci-lint
RUN wget https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh -P /tmp \
    && sh /tmp/install.sh -b /usr/local/bin v1.56.2 \
    && rm /tmp/install.sh

# Clean up
RUN apt-get remove -y wget \
    && apt-get clean \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /cms-and-go

# Copy the source code
COPY . .

# Add health check (optional)
HEALTHCHECK CMD go version || exit 1

# Expose necessary ports (if any)
EXPOSE 8080
