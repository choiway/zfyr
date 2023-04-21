package web

import (
	"embed"
	"github.com/labstack/echo/v4"
)

var (
    //go:generate npm i
    //go:generate npm run build
	//go:embed all:dist
	dist embed.FS
	//go:embed dist/index.html
	indexHTML embed.FS

	distDirFS     = echo.MustSubFS(dist, "dist")
	distIndexHtml = echo.MustSubFS(indexHTML, "dist")
)

func RegisterHandlers(e *echo.Echo) {
    e.StaticFS("/", distDirFS)
	e.FileFS("/", "index.html", distIndexHtml)
	e.FileFS("/shares", "index.html", distIndexHtml)
}
