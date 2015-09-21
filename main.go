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

	// Group using gin.BasicAuth() middleware
	authorized := r.Group("/", gin.BasicAuth(config.Accounts))

	for _, f := range config.StaticFS {
		authorized.StaticFS("/"+f, http.Dir(f))
	}

	authorized.GET(config.HomeUrl, func(c *gin.Context) {
		c.JSON(http.StatusOK, "home!")
	})

	r.Run(config.ListenAddress)
}
