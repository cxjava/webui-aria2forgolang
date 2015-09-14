package main

import (
	"fmt"
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

	for _, f := range config.StaticFS {
		r.StaticFS("/"+f, http.Dir(f))
	}

	r.GET(config.HomeUrl, func(c *gin.Context) {
		c.JSON(http.StatusOK, "home!")
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
