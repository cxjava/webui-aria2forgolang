package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func Hello(c *echo.Context) error {
	return c.Render(http.StatusOK, "1/index.tmpl", "1111")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	e.SetRenderer(DefaultHTMLBinData())
	e.Get("/1", Hello)

	e.Get("/2", func(c *echo.Context) error {
		return c.Render(http.StatusOK, "2/index.tmpl", "2222")
	})

	// Start server
	e.Run(":1323")
}
