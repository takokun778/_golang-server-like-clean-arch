#!/bin/sh -eu

export ENV=local
export DB_NAME=template
export DB_USER=template
export DB_PASS=template
export DB_PORT=5432
export DEBUG=ON

go run ./app/main.go
