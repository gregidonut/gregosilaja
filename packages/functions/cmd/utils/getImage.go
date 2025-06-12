package utils

import (
	"bytes"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	_ "image/jpeg"
)

func GetImageDimensions(imageData []byte) (width, height int, err error) {
	reader := bytes.NewReader(imageData)
	img, _, err := image.Decode(reader)
	if err != nil {
		return 0, 0, err
	}
	bounds := img.Bounds()
	return bounds.Dx(), bounds.Dy(), nil
}

func ResConv(imageData []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(imageData))
	if err != nil {
		return nil, err
	}

	resizedImg := imaging.Resize(img, 400, 0, imaging.Lanczos)

	var webpBuf bytes.Buffer
	err = webp.Encode(&webpBuf, resizedImg, &webp.Options{
		Lossless: false,
		Quality:  75,
	})
	if err != nil {
		return nil, err
	}

	return webpBuf.Bytes(), nil
}
