package main

import (
	"html/template"
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

	html := template.Must(template.New("").Delims("[[", "]]").ParseFiles("./views/index.html"))
	r.SetHTMLTemplate(html)

	r.Static("/assets", "./assets")

	r.GET(config.HomeUrl, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(config.ListenAddress)
}
