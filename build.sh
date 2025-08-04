#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o github_dlcount_linux -ldflags="-s -w" -trimpath ./cmd/github_dlcount

GOOS=windows GOARCH=amd64 go build -o github_dlcount_windows.exe -ldflags="-s -w" -trimpath ./cmd/github_dlcount
