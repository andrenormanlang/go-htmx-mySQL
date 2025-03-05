#! /usr/bin/env bash

set -euo pipefail

git config --global --add safe.directory '*'

cd /cmsgo/migrations
GOOSE_DRIVER="mysql" GOOSE_DBSTRING="root:root@tcp(mariadb:3306)/cms-and-go" goose up

cd /cmsgo

# Check if we're running in admin mode
if [ "$1" = "--admin" ]; then
  # Run admin app with admin-specific air config
  air -c ./docker/air-admin.toml
else
  # Run regular app
  air -c ./docker/air.toml
fi