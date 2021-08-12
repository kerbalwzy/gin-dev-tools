package kerbalwzygo

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
)

func GenerateQrCode(content string, width, height int) ([]byte, error) {
	res := new(bytes.Buffer)
	qrCode, err := qr.Encode(content, qr.H, qr.Auto)
	if nil != err {
		return nil, err
	}
	qrCode, err = barcode.Scale(qrCode, width, height)
	err = png.Encode(res, qrCode)
	if nil != err {
		return nil, err
	}
	return res.Bytes(), nil
}
