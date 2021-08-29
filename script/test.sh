#!/bin/sh -eu

export ENV=test
export DB_NAME=${ENV}
export DB_USER=${ENV}
export DB_PASS=${ENV}
export DB_PORT=54321

go test -v ./...
