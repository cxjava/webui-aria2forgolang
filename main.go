package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusOK, "home!")
	})

	r.Run(config.ListenAddress)
}
