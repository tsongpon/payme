package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tsongpon/payme/handler"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/qrcode", handler.CreateQRCode)
	// e.GET("/healthz", func(c echo.Context) error { return c.String(200, "ok") })
	e.Logger.Fatal(e.Start(":8080"))
}
