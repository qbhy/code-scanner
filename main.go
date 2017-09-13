package main

import (
	"os"
	"image/png"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode"
)

func main() {
	// Create the barcode
	qrCode, _ := qr.Encode("测试一下呗", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}
