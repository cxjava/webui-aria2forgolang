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

	authorized := r.Group("/")   

	// Group using gin.BasicAuth() middleware
	if len(config.Accounts)>0{     
		authorized.Use(gin.BasicAuth(config.Accounts))
	}
	
	for _, f := range config.StaticFS {
		authorized.StaticFS("/"+f, http.Dir(f))
	}

	authorized.GET(config.HomeUrl, func(c *gin.Context) {
		c.JSON(http.StatusOK, "welcome home!")
	})

	r.GET(config.HomeUrl+"/ping", func(c *gin.Context) {
		addressAndPort := c.Request.RemoteAddr
		fmt.Println("Remote address:", addressAndPort)
		remoteAddress = addressAndPort[:strings.Index(addressAndPort, ":")]
		c.JSON(http.StatusOK, "pong!")
	})

	authorized.GET(config.HomeUrl+"/ip", func(c *gin.Context) {
		c.JSON(http.StatusOK, remoteAddress)
	})

	r.Run(config.ListenAddress)
}
