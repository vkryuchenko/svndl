#!/bin/bash
go get -u bitbucket.org/slavyan85/svnwrapper
export GOPATH=$GOPATH:`pwd`

go build -o SvnDataLoader src\svndataloader.go

export GOOS=windows
export GOARCH=amd64
go build -o SvnDataLoader.exe src\svndataloader.go
