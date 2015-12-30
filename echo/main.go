package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var (
	address = flag.String("a", ":18080", `bind address.`)
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

	fmt.Println("address " + *address)
	e.Run(*address)
}
