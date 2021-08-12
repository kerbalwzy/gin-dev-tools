package kerbalwzygo

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestGenerateQrCode(t *testing.T) {
	content := "hello word"
	width, height := 256, 256
	data, err := GenerateQrCode(content, width, height)
	if nil != err {
		log.Fatal(err)
	}

	ioutil.WriteFile("test_qr_code.png", data, 0777)
}
