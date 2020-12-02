package ifile

import (
	"os"
)

// IFileOperator はfile.goを操作するインターフェース。
type IFileOperator interface {
	// Find は条件を指定してファイル一覧を取得する。
	Find(params *FindFileParameters) ([]string, error)

	// Create は画像を書き出すためのファイルを生成する。
	Create(
		path string,
		file *os.File,
		ext string,
	) (*os.File, error)
}

// FindFileParameters はファイルを検索するための条件。
type FindFileParameters struct {
	// ディレクトリ。
	Dir string
	// 拡張子。
	Ext string
}
