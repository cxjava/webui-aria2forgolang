package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

var remoteAddress = ""

func init() {
	readConfig()
}

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	for _, f := range config.StaticFS {
		r.StaticFS("/"+f, http.Dir(f))
	}

	r.GET(config.HomeUrl, func(c *gin.Context) {
		c.JSON(http.StatusOK, "welcome home!")
	})

	r.GET(config.HomeUrl+"/ping", func(c *gin.Context) {
		addressAndPort := c.Request.RemoteAddr
		fmt.Println("Remote address:", addressAndPort)
		remoteAddress = addressAndPort[:strings.Index(addressAndPort, ":")]
		c.JSON(http.StatusOK, "pong!")
	})

	r.GET(config.HomeUrl+"/aria2", func(c *gin.Context) {
		c.Redirect(http.StatusFound, remoteAddress+config.ForwardPort)
	})

	r.Run(config.ListenAddress)
}
