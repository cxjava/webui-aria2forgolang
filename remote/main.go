package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/kataras/iris"
)

var (
	address       = flag.String("a", ":15555", `bind address.`)
	remoteAddress = ""
)

func init() {
	flag.Parse()
}

func main() {

	iris.Get("/ping", func(c *iris.Context) {
		remoteAddress = c.RequestIP()
		fmt.Println("Remote address:", remoteAddress)
		c.JSON(http.StatusOK, "pong!")
		c.ServeFile("./myfolder/staticfile.txt")
	})
	iris.Get("/ip", func(c *iris.Context) {
		c.WriteHTML(http.StatusOK, `<html>
			<body>
			<a href="http://`+remoteAddress+`:2987/" target="_blank">home pi</a>
			</body>
			</html>`)
	})
	iris.Get("/aria/*file", func(ctx *iris.Context) {
		filepath := ctx.Param("file")
		if len(filepath) == 0 {
			filepath = "/"
		}
		fmt.Println(filepath)
		if len(filepath) == 1 {
			filepath += "index.html"
		}
		ff, err := aria.Open(filepath)
		if err != nil {
			fmt.Println(err)
			ctx.NotFound()
			return
		}
		fmt.Println(ff)
		ctx.ServeContent(ff, ff.(os.FileInfo).Name(), ff.(os.FileInfo).ModTime())
	})
	// pathSeparator := string(os.PathSeparator)
	iris.Get("/yaaw/*file", func(ctx *iris.Context) {
		filepath := ctx.Param("file")
		if len(filepath) == 0 {
			filepath = "/"
		}
		fmt.Println(filepath)
		if len(filepath) == 1 {
			filepath += "index.html"
		}

		ff, err := yaaw.Open(filepath)
		if err != nil {
			fmt.Println(err)
			ctx.NotFound()
			return
		}
		fmt.Println(ff)
		ctx.ServeContent(ff, ff.(os.FileInfo).Name(), ff.(os.FileInfo).ModTime())
	})
	fmt.Println("address " + *address)
	iris.Listen(*address)
}
