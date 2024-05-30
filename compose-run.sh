#! /usr/bin/env bash

set -euo pipefail

git config --global --add safe.directory '*'

cd /gocms/migrations
GOOSE_DRIVER="mysql" GOOSE_DBSTRING="root:shiva7@tcp(mariadb:3306)/cms-and-go" goose up

cd /gocms
air