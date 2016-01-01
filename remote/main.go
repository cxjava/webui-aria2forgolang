package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var (
	address       = flag.String("a", ":15555", `bind address.`)
	remoteAddress = ""
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

	e.Get("/ping", func(c *echo.Context) error {
		addressAndPort := c.Request().RemoteAddr
		fmt.Println("Remote address:", addressAndPort)
		remoteAddress = addressAndPort[:strings.Index(addressAndPort, ":")]
		return c.JSON(http.StatusOK, "pong!")
	})

	e.Get("/ip", func(c *echo.Context) error {
		return c.HTML(http.StatusOK, `<html>
			<body>
			<a href="http://`+remoteAddress+`:2987/" target="_blank">home pi</a>
			</body>
			</html>`)
	})

	fmt.Println("address " + *address)
	e.Run(*address)
}
