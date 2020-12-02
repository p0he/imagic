package file

import (
	"os"
	"path/filepath"
	"strings"

	"imagic/src/file/ifile"
)

// fileOperator はIFileOperatorの実装である。
type fileOperator struct {
	operator ifile.IFileOperator
}

// NewFileOperator はfileOperatorのインスタンスを生成する。
func NewFileOperator() *fileOperator {
	return &fileOperator{}
}

// Find はIFileOperator.Findの実装である。
func (f *fileOperator) Find(
	params *ifile.FindFileParameters,
) ([]string, error) {
	var paths []string
	err := filepath.Walk(params.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == params.Ext {
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}

// Create はIFileOperator.Createの実装である。
func (f *fileOperator) Create(
	path string,
	file *os.File,
	ext string,
) (*os.File, error) {
	outPath := f.generateOutPath(path, ext)
	file, err := os.Create(outPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *fileOperator) generateOutPath(
	path string,
	ext string,
) string {
	fromExt := filepath.Ext(path)
	return strings.TrimSuffix(path, fromExt) + ext
}
