#!/bin/bash

# Build the executable
cd ..
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/UpdateDomainRecord
cd build || return
