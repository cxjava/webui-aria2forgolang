package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var (
	address       = flag.String("a", ":18080", `bind address.`)
	remoteAddress = flag.String("ra", "http://54.169.232.135:15555/ping", `remote address.`)
	pingInterval  = flag.Int("p", 5, `Ping Interval.`)
)

func init() {
	flag.Parse()
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	yaawHandler := http.StripPrefix("/yaaw/", http.FileServer(yaaw))
	e.Get("/yaaw/*", func(c *echo.Context) error {
		yaawHandler.ServeHTTP(c.Response().Writer(), c.Request())
		return nil
	})

	ariaHandler := http.StripPrefix("/aria/", http.FileServer(aria))
	e.Get("/aria/*", func(c *echo.Context) error {
		ariaHandler.ServeHTTP(c.Response().Writer(), c.Request())
		return nil
	})

	Tick()

	fmt.Println("address " + *address)
	e.Run(*address)
}

func Tick() {
	go func() {
		hb := time.Tick(time.Duration(*pingInterval) * time.Minute)
		for {
			select {
			case <-hb:
				ping()
			}
		}
	}()
}

func ping() {
	resp, err := http.Get(*remoteAddress)
	if err != nil {
		fmt.Println("http.Get err:", err)
	}
	defer resp.Body.Close()
	fmt.Println("StatusCode:", resp.StatusCode)
}
