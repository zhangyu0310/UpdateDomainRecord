#!/bin/bash

# Build the executable
cd ..
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o ./bin/UpdateDomainRecord
cd build || return
