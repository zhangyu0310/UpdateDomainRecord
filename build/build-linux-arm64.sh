#!/bin/bash

# Build the executable
cd ..
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/UpdateDomainRecord
cd build || return
