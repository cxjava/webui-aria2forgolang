package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.LoadHTMLGlob("templates/**/*")
	router.HTMLRender = DefaultHTMLBinData()

	// router.LoadHTMLFiles("templates/2/index.tmpl", "templates/1/index.tmpl")
	router.GET("/2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "2/index.tmpl", gin.H{
			"title": "222 website",
		})
	})
	router.GET("/1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "1/index.tmpl", gin.H{
			"title": "111 website",
		})
	})
	router.Run(":8080")
}
