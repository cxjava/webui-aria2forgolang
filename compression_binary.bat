go build -ldflags="-s -w" -o="compression_binary.exe"
upx --brute compression_binary.exe