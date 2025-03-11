#!/bin/bash

# Run docker-compose from the docker directory
cd "$(dirname "$0")"
docker-compose up "$@" 