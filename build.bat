go get -u bitbucket.org/slavyan85/svnwrapper
set GOPATH=%GOPATH%;%~dp0

go build -o SvnDataLoader.exe src\svndataloader.go

set GOOS=linux
set GOARCH=amd64
go build -o SvnDataLoader src\svndataloader.go
