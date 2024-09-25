#! /usr/bin/env bash

set -euo pipefail

# Navigate to the application directory
cd /cmsgo

# Run any necessary migrations, set up the environment, etc.
GOOSE_DRIVER="mysql" GOOSE_DBSTRING="root:root@tcp(mariadb:3306)/cms-and-go" goose up

# Start your application (replace with your actual command)
air -c ./docker/.admin.air.toml
