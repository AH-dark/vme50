#!/usr/bin/bash

export CGO_ENABLED=0

export GOOS=linux GOARCH=amd64
go build -a -o build/linux_amd64_random_donate .

export GOOS=linux GOARCH=arm64
go build -a -o build/linux_arm64_random_donate .

export GOOS=darwin GOARCH=amd64
go build -a -o build/darwin_amd64_random_donate .

export GOOS=windows GOARCH=amd64
go build -a -o build/windows_amd64_random_donate.exe .
