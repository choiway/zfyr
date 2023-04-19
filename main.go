package main

import (
	"github.com/labstack/echo/v4"
    "zyfr/web"
	"net/http"
)

func main() {
	e := echo.New()

    web.RegisterHandlers(e)

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Peeeps")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
