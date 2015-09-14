package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

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

	Tick()

	r.Run(config.ListenAddress)
}

func Tick() {
	go func() {
		hb := time.Tick(time.Duration(config.PingInterval) * time.Minute)
		for {
			select {
			case <-hb:
				ping()
			}
		}
	}()
}

func ping() {
	resp, err := http.Get(config.RemoteAddress)
	if err != nil {
		fmt.Println("http.Get err:", err)
	}
	defer resp.Body.Close()
	fmt.Println("StatusCode:", resp.StatusCode)
}
