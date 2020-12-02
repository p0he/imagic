package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"


	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"

	"imagic/src/converter/iconverter"
)

// converterOperator はIConverterOperatorの実装である。
type converterOperator struct {
	operator iconverter.ConverterOperator
}

// NewConverterOperator はconverterOperatorのインスタンスを生成する。
func NewConverterOperator() *converterOperator {
	return &converterOperator{}
}

// Convert はIConverter.ConvertImageの実装である。
func (c *converterOperator) ConvertImage(
	fromFile *os.File,
	toFile io.Writer,
	toExt string,
) error {
	img, _, err := image.Decode(fromFile)
	if err != nil {
		return err
	}

	switch toExt {
	case ".jpg", ".jpeg":
		if err := jpeg.Encode(toFile, img, nil); err != nil {
			return err
		}

	case ".png":
		if err := png.Encode(toFile, img); err != nil {
			return err
		}

	case ".gif":
		if err := gif.Encode(toFile, img, nil); err != nil {
			return err
		}

	case ".tiff":
		if err := tiff.Encode(toFile, img, nil); err != nil {
			return err
		}

	case ".bmp":
		if err := bmp.Encode(toFile, img); err != nil {
			return err
		}
	}

	return nil
}
