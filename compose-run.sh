#! /usr/bin/env bash

set -euo pipefail

git config --global --add safe.directory '*'

cd /CMSGO/migrations
GOOSE_DRIVER="mysql" GOOSE_DBSTRING="root:root@tcp(mariadb:3306)/cms-and-go" goose up

cd /CMSGO
air