package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/soma/:x/:y", Sum)
	e.Logger.Print("Listening on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
