package cmd

import (
	"flag"
	"fmt"
	"os"

	"imagic/src/converter/iconverter"
	"imagic/src/file/ifile"
	"imagic/src/validator/ivalidator"
)

var (
	fromExt string
	toExt   string
)

func init() {
	flag.StringVar(&fromExt, "f", ".jpg", "変換元の画像ファイルの拡張子（例: .jpg")
	flag.StringVar(&toExt, "t", ".png", "変換先の画像ファイルの拡張子（例: .png")
}

// cmdOperator はICMDOperatorの実装である。
type cmdOperator struct {
	fileOperator      ifile.IFileOperator
	convertOperator   iconverter.ConverterOperator
	validatorOperator ivalidator.ValidatorOperator
}

// NewCMDOperator は cmdOperatorのインスタンスを生成する。
func NewCMDOperator(
	fileOperator ifile.IFileOperator,
	convertOperator iconverter.ConverterOperator,
	validatorOperator ivalidator.ValidatorOperator,
) *cmdOperator {
	return &cmdOperator{
		fileOperator:      fileOperator,
		convertOperator:   convertOperator,
		validatorOperator: validatorOperator,
	}
}

// Run はICMDOperator.Runの実装である。
func (c *cmdOperator) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please specify directory")
	}

	err := c.validatorOperator.ValidateExt(toExt)
	if err != nil {
		return err
	}

	files, err := c.fileOperator.Find(&ifile.FindFileParameters{
		Dir: args[0],
		Ext: fromExt,
	})
	if err != nil {
		return fmt.Errorf("failed to load files")
	}
	if len(files) == 0 {
		fmt.Println("failed to fined specified files")
		os.Exit(1)
	}

	for _, file := range files {
		fromFile, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("failed to open %s", file)
		}
		defer fromFile.Close()

		toFile, err := c.fileOperator.Create(file, fromFile, toExt)
		if err != nil {
			return fmt.Errorf("failed to create %s for output %s", file, toExt)
		}
		defer toFile.Close()

		if err := c.convertOperator.ConvertImage(fromFile, toFile, toExt); err != nil {
			if err != nil {
				os.Remove(toFile.Name())
				return fmt.Errorf("failed to convert %s", fromFile.Name())
			}
		}
	}

	fmt.Println("conversion successfully")
	return nil
}
