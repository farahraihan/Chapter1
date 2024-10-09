package main

import (
	"chapter1/internal/factory"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	factory.InitFactory(e)
	e.Logger.Error(e.Start(":8000"))
}
