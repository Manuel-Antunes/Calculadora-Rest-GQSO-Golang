package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Result struct {
	Value float64 `json:"value" bson:"value,omitempty"`
}

func Sum(c echo.Context) error {
	x, err := strconv.ParseFloat(c.Param("x"), 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Parametro x inválido")
	}
	y, err := strconv.ParseFloat(c.Param("y"), 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Parametro y inválido")
	}
	return c.JSON(http.StatusOK, Result{Value: x + y})
}

func Sub(c echo.Context) error {
	x, err := strconv.ParseFloat(c.Param("x"), 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Parametro x inválido")
	}
	y, err := strconv.ParseFloat(c.Param("y"), 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Parametro y inválido")
	}
	return c.JSON(http.StatusOK, Result{Value: x - y})
}
