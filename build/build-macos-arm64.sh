#!/bin/bash

# Build the executable
cd ..
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./bin/UpdateDomainRecord
cd build || return
