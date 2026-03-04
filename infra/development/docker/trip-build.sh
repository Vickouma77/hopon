#!/bin/bash
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o build/trip-service ./services/trip-service/cmd/main.go
