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
	out, err := exec.Command("zpool", "status").CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("Zpool status:")
	fmt.Println(string(out))
	e := echo.New()

	web.RegisterHandlers(e)

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Salutations, Peeeps")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

