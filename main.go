package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os/exec"
	"zyfr/web"
)

func main() {
	out, err := exec.Command("ls", "-l").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
	e := echo.New()

	web.RegisterHandlers(e)

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Peeeps")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
