package routes

import (

	"github.com/labstack/echo/v4"

	"echo-template/home"

)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return home.HomeHandler(c)
	})
}
