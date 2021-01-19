#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
mv main evm-server
chmod 777 evm-server
mv evm-server ~/fsdownload/