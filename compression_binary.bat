go build -ldflags="-s -w" -o="compression_binary.exe"
upx --best compression_binary.exe
pause