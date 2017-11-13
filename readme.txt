// frist download yaaw-master.zip and webui-aria2-master.zip and compiler.jar
// delete folder webui-aria2-master\flags\ ;  webui-aria2-master\screenshots\  国际化文件 style.css里面第一行
// and run compilerJS.bat
// open yaaw_vfsdata.go  and delete _vfsgen_fs

.\yaaw_vfsdata.go:8:2: imported and not used: "bytes"
.\yaaw_vfsdata.go:9:2: imported and not used: "compress/gzip"
.\yaaw_vfsdata.go:10:2: imported and not used: "fmt"
.\yaaw_vfsdata.go:11:2: imported and not used: "io"
.\yaaw_vfsdata.go:12:2: imported and not used: "io/ioutil"
.\yaaw_vfsdata.go:15:2: imported and not used: "path" as pathpkg

go generate

 go build -tags vfs -i -ldflags "-w -s"
 
 upx --brute 