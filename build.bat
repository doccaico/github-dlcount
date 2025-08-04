@echo off

setlocal

set GOOS=windows
set GOARCH=amd64

go build -o github_dlcount_windows.exe -ldflags="-s -w" -trimpath .\cmd\github-dlcount

set GOOS=linux
set GOARCH=amd64

go build -o github_dlcount_linux -ldflags="-s -w" -trimpath .\cmd\github-dlcount
