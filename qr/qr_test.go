package qr_test

import (
	"testing"

	"github.com/tsongpon/payme/qr"
)

// func TestCreateQRCode(t *testing.T) {
// 	qrCode := qr.CreateQRCode("0899999999", 20)
// 	img := qr.CraeteQRImage(qrCode)
// 	//save the imgByte to file
// 	out, err := os.Create("./QRImg666.png")

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	err = png.Encode(out, img)

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	if qrCode != "00020101021229370016A000000677010111011300668999999995802TH5303764540520.0063044DAF" {
// 		t.Error("qr.Create(\"0899999999\", \"20\") should return 00020101021229370016A000000677010111011300668999999995802TH5303764540520.0063044DAF, but got ", qrCode)
// 	}
// }

func TestCreatePrompPayQRValidPhoneNumber(t *testing.T) {
	expected := "00020101021129370016A000000677010111011300660000000005802TH530376463048956"
	qr := qr.CreatePrompPayQRCode("0000000000", 0)
	if qr != "00020101021129370016A000000677010111011300660000000005802TH530376463048956" {
		t.Error("CreatePrompPayQR should return " + expected + ", but got " + qr)
	}
}

func TestCreatePrompPayQRValidCitizenNumber(t *testing.T) {
	expected := "00020101021129370016A000000677010111011331305551294625802TH530376463045858"
	qr := qr.CreatePrompPayQRCode("3130555129462", 0)
	if qr != "00020101021129370016A000000677010111011331305551294625802TH530376463045858" {
		t.Error("CreatePrompPayQR(\"3130555129462\", 0) should return " + expected + ", but got " + qr)
	}
}

func TestQRCodeToImage(t *testing.T) {
	qrCode := qr.CreatePrompPayQRCode("0809710099", 50)
	qr.QRCodeToImage(qrCode)
}
