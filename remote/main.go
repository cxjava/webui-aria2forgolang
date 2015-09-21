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

	// Group using gin.BasicAuth() middleware
	authorized := r.Group("/", gin.BasicAuth(config.Accounts))

	for _, f := range config.StaticFS {
		authorized.StaticFS("/"+f, http.Dir(f))
	}

	authorized.GET(config.HomeUrl, func(c *gin.Context) {
		c.JSON(http.StatusOK, "welcome home!")
	})

	authorized.GET(config.HomeUrl+"/ping", func(c *gin.Context) {
		addressAndPort := c.Request.RemoteAddr
		fmt.Println("Remote address:", addressAndPort)
		remoteAddress = addressAndPort[:strings.Index(addressAndPort, ":")]
		c.JSON(http.StatusOK, "pong!")
	})

	authorized.GET(config.HomeUrl+"/aria2", func(c *gin.Context) {
		c.Redirect(http.StatusFound, remoteAddress+config.ForwardPort)
	})

	r.Run(config.ListenAddress)
}
