package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsongpon/payme/qr"
	"github.com/tsongpon/payme/transport"
)

func CreateQRCode(c echo.Context) error {
	r := new(transport.PrompPayQRCodeRequest)
	if err := c.Bind(r); err != nil {
		return err
	}
	qrCode := qr.CreatePrompPayQRCode(r.Target, r.Amount)
	data := qr.QRCodeToImage(qrCode)
	return c.Blob(http.StatusOK, "image/png", data)
}
