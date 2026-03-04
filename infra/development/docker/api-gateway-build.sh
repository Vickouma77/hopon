#!/bin/bash
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o build/api-gateway ./services/api-gateway
