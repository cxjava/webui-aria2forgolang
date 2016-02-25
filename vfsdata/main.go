package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/SimonWaldherr/golibs/file"
	"github.com/Workiva/go-datastructures/queue"
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
	var wg sync.WaitGroup
	rb := queue.NewRingBuffer(4)

	go func() {
		for {
			filepath, err := rb.Get()
			if err != nil {
				fmt.Println("rb.Get:", err)
			}
			fmt.Println(filepath)
			cmd := exec.Command("java", "-jar", "compiler.jar", `--js_output_file="`+filepath.(string)+`"`, filepath.(string))
			data, errr := cmd.Output()
			if errr != nil {
				fmt.Println("cmd.Output:", errr)
			}
			fmt.Println(string(data))
			wg.Done()
		}
	}()

	err := file.Each("..", true, func(filename, extension, filepath string, dir bool, fileinfo os.FileInfo) {
		if extension == "js" && !dir {
			wg.Add(1)
			err := rb.Put(filepath)
			if err != nil {
				fmt.Println("rb.Put:", err)
			}
			fmt.Println(filename)
		}
	})
	if err != nil {
		log.Fatalln(err)
	}
	wg.Wait()
	rb.Dispose()
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
