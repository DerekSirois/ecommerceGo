package handler

import "github.com/labstack/echo/v4"

func Index(c echo.Context) error {
	return c.Render(200, "index.html", nil)
}

func RegisterPage(c echo.Context) error {
	return c.Render(200, "register.html", nil)
}

func LoginPage(c echo.Context) error {
	return c.Render(200, "login.html", nil)
}
