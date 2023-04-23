package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"net"
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

	// Get Ip address

	fmt.Println("IP Address:")
	ip := GetOutboundIP()
	fmt.Println(ip)

	e := echo.New()

	// Static Files

	web.RegisterHandlers(e)

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	// HTML Files

	// e.File("/", "index.html")
	e.Renderer = t
	e.GET("/", Index)
	e.File("/shares", "index.html")
	e.File("/about", "index.html")

	// API Endpoints

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Salutations, Peeeps")
	})

	e.Logger.Fatal(e.Start(":1323"))

}

type IndexPayload struct {
    Title string
}

var indexPayload = IndexPayload{
    Title: "Homies, this is from index payload",
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", indexPayload)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
