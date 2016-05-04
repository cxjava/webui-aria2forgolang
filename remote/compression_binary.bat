set GOARCH=amd64
set GOOS=linux
go build -ldflags="-s -w" -o="remote"
..\upx --brute remote
pause