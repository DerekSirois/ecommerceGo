package server

import (
	"ecommerceGo/internal/handler"

	"github.com/labstack/echo/v4"
)

func initRoutes(e *echo.Echo) {
	e.GET("/", handler.Index)
}
