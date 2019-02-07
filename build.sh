#!/bin/bash
export GOPATH=$GOPATH:$(pwd)
go build -o svndl src/svndataloader.go

export GOOS=windows
export GOARCH=amd64
go build -o svndl.exe src/svndataloader.go
