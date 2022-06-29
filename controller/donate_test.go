package controller

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"image"
	_ "image/png"
	"os"
	"testing"
)

func TestImageDecode(t *testing.T) {
	asserts := assert.New(t)
	file, err := os.ReadFile("test_qrcode.png")
	asserts.NoError(err)
	imgReader := bytes.NewBuffer(file)
	img, _, err := image.Decode(imgReader)
	asserts.NoError(err)
	asserts.NotNil(img)
}
