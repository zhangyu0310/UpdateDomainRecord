#!/bin/bash

# Build the executable
cd ..
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/UpdateDomainRecord
cd build || return
