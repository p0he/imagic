package iconverter

import (
	"io"
	"os"
)

// ConverterOperator はconverter.goを操作するインターフェース。
type ConverterOperator interface {
	// ConvertImage は画像ファイルを指定された拡張子に変換する。
	// 第一引数には変換元の画像ファイルを指定する。
	// 第二引数には変換先の画像ファイルを書き出すファイルを指定する。
	ConvertImage(
		fromFile *os.File,
		toFile io.Writer,
		toExt string,
	) error
}
