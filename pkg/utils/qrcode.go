package utils

import (
	"bytes"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	"io"
	"mime/multipart"
)

func ParseQRCode(QRCode *multipart.FileHeader) (string, error) {
	fileReader, err := QRCode.Open()
	imgReader := bytes.NewBuffer(nil)
	_, err = io.Copy(imgReader, fileReader)
	if err != nil {
		return "", err
	}

	// TODO Handle Error
	// prepare BinaryBitmap
	img, _, err := image.Decode(imgReader)
	if err != nil {
		return "", err
	}
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", err
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}

	text := result.GetText()
	Log().Debug("解析二维码成功 %s", text)

	return text, nil
}
