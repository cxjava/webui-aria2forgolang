go build -ldflags="-s -w" -o="compression_binary.exe"
..\upx -f --brute compression_binary.exe
pause