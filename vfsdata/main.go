package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/SimonWaldherr/golibs/file"
	"github.com/shurcooL/vfsgen"
)

// frist download yaaw-master.zip and webui-aria2-master.zip and compiler.jar
// delete folder webui-aria2-master\flags\ ;  webui-aria2-master\screenshots\
// and run compilerJS.bat
// open yaaw_vfsdata.go  and delete var _vfsgen_fs
func main() {
	// Compress()
	vfsgenFile()
}
func Compress() {
	var w sync.WaitGroup
	err := file.Each("..", true, func(filename, extension, filepath string, dir bool, fileinfo os.FileInfo) {
		if extension == "js" && !dir {
			fmt.Println(filename)
			w.Add(1)
			go func() {
				cmd := exec.Command("java", "-jar", "compiler.jar", `--js_output_file="`+filepath+`"`, `"`+filepath+`"`)
				data, err := cmd.Output()
				if err != nil {
					fmt.Println("java fail:", err)
				}
				fmt.Println(string(data))
				w.Done()
			}()
		}
	})
	if err != nil {
		log.Fatalln(err)
	}
	w.Wait()
}
func vfsgenFile() {
	var fs http.FileSystem = http.Dir("webui-aria2-master")

	err := vfsgen.Generate(fs, vfsgen.Options{
		VariableName: "aria",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fs = http.Dir("yaaw-master")

	err = vfsgen.Generate(fs, vfsgen.Options{
		VariableName: "yaaw",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
