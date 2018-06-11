// frist download yaaw-master.zip and webui-aria2-master.zip and compiler.jar
// delete folder webui-aria2-master\flags\ ;  webui-aria2-master\screenshots\ 
运行del.bat
 
  style.css里面第一行 
init.js 里面的语言文件
remove .translations('id_ID', mergeTranslation(translations.id_ID, translations.en_US))

不能压缩这个函数  webui.config(function ($translateProvider, $locationProvider) {
  $translateProvider
  
index.html 里面的语言文件引用，还有菜单
<script src="js/translate/id_ID.js"></script>
<li ng-click="changeLanguage('id_ID')"
删除js国际化文件
// and run compilerJS.bat
// open yaaw_vfsdata.go  and delete _vfsgen_fs vfsgen۰FS

.\yaaw_vfsdata.go:8:2: imported and not used: "bytes"
.\yaaw_vfsdata.go:9:2: imported and not used: "compress/gzip"
.\yaaw_vfsdata.go:10:2: imported and not used: "fmt"
.\yaaw_vfsdata.go:11:2: imported and not used: "io"
.\yaaw_vfsdata.go:12:2: imported and not used: "io/ioutil"
.\yaaw_vfsdata.go:15:2: imported and not used: "path" as pathpkg

go generate

go build -tags vfs -i -ldflags "-w -s"
 
go build -tags release -i -ldflags "-w -s"
 
 upx --brute aria2.exe
 
IntelliJ IDEA Evaluator