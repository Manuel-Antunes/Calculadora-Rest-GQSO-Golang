package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/soma/:x/:y", sumHandler)
	e.Logger.Print("Listening on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}

func sumHandler(c echo.Context) error {
	op1 := c.Param("x")
	op2 := c.Param("y")
	result, err := Sum(op1, op2)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("sum error:%q", err))
	}
	return c.JSON(http.StatusOK, *result)
}
