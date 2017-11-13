// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = http.Dir("vfsdata/webui-aria2-master")

	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "main",
		BuildTags:    "vfs",
		VariableName: "aria",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fs = http.Dir("vfsdata/yaaw-master")

	err = vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "main",
		BuildTags:    "vfs",
		VariableName: "yaaw",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
