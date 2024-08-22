package qr_test

import (
	"fmt"
	"testing"

	"github.com/tsongpon/payme/internal/qr"
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

func TestChecksum(t *testing.T) {
	crc := qr.CalChecksum("00020101021129370016A000000677010111011300660000000005802TH53037646304")
	if crc != "0x8956" {
		t.Error("checksum should return 0x8956, but got ", crc)
	}
}

func TestCreatePrompPayQR(t *testing.T) {
	qr := qr.CreatePrompPayQR()
	fmt.Println(qr)
}
