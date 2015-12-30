package main

import (
	"flag"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	address = flag.String("a", ":18080", `bind address.`)
)

func init() {
	flag.Parse()
}

func main() {
	r := gin.Default()

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.StaticFS("/yaaw", yaaw)
	r.StaticFS("/aria", aria)

	r.Run(*address)
}
