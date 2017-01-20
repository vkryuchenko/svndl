set GOPATH=%GOPATH%;%~dp0

go build -o svndl.exe src\svndataloader.go

set GOOS=linux
set GOARCH=amd64
go build -o svndl src\svndataloader.go
