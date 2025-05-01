package qr_test

import (
	"testing"

	"github.com/tsongpon/payme/qr"
)

func TestCreatePrompPayQRValidPhoneNumber(t *testing.T) {
	expected := "00020101021129370016A000000677010111011300660000000005802TH530376463048956"
	qr, err := qr.CreatePrompPayQRCode("0000000000", 0)
	if err != nil {
		t.Error("CreatePrompPayQR(\"0000000000\", 0) should not return error")
	}
	if qr != "00020101021129370016A000000677010111011300660000000005802TH530376463048956" {
		t.Error("CreatePrompPayQR(\"0000000000\", 0) should return " + expected + ", but got " + qr)
	}
}

func TestCreatePrompPayQRValidCitizenNumber(t *testing.T) {
	expected := "00020101021129370016A000000677010111011331305551294625802TH530376463045858"
	qr, err := qr.CreatePrompPayQRCode("3130555129462", 0)
	if err != nil {
		t.Error("CreatePrompPayQR(\"3130555129462\", 0) should not return error")
	}
	if qr != "00020101021129370016A000000677010111011331305551294625802TH530376463045858" {
		t.Error("CreatePrompPayQR(\"3130555129462\", 0) should return " + expected + ", but got " + qr)
	}
}

func TestQRCodeToImage(t *testing.T) {
	qrCode, err := qr.CreatePrompPayQRCode("0809710099", 50)
	if err != nil {
		t.Error("CreatePrompPayQR(\"0809710099\", 0) should not return error")
	}
	qr.QRCodeToImage(qrCode)
}
