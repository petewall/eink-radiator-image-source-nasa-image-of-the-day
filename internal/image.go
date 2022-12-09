package internal

import (
	"image"
	"image/png"
	"io"
	"os"

	imageConf "github.com/petewall/eink-radiator-image-source-image/pkg"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . ImageProcessor
type ImageProcessor func(url string, width, height int) (image.Image, error)

var ProcessImage ImageProcessor = func(url string, width, height int) (image.Image, error) {
	imageConfig := &imageConf.Config{
		Source: url,
		Scale:  imageConf.ScaleCover,
	}
	return imageConfig.GenerateImage(width, height)
}

//counterfeiter:generate . ImageEncoder
type ImageEncoder func(w io.Writer, i image.Image) error

var EncodeImage ImageEncoder = png.Encode

//counterfeiter:generate . ImageWriter
type ImageWriter func(file string, i image.Image) error

var WriteImage ImageWriter = func(file string, i image.Image) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	err = EncodeImage(f, i)
	if err != nil {
		return err
	}
	return f.Close()
}
