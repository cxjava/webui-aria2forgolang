package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
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


	e.Get("/yaaw/*", func(c echo.Context) error {
		yaawHandler.ServeHTTP(c.Response().(*standard.Response).ResponseWriter, c.Request().(*standard.Request).Request)
		return nil
	})

	ariaHandler := http.StripPrefix("/aria/", http.FileServer(aria))
	e.Get("/aria/*", func(c echo.Context) error {
		ariaHandler.ServeHTTP(c.Response().(*standard.Response).ResponseWriter, c.Request().(*standard.Request).Request)
		return nil
	})


	fmt.Println("address " + *address)
	e.Run(standard.New(*address))
}
