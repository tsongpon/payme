package main

import (
	"github.com/labstack/echo/v4"
	"github.com/tsongpon/payme/handler"
)

func main() {
	e := echo.New()
	e.POST("/qrcode", handler.CreateQRCode)
	e.Logger.Fatal(e.Start(":1323"))
}
