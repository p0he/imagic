package validator

import (
	"fmt"
	"strings"

	"imagic/src/validator/ivalidator"
)

type FilNameExtension int

const (
	JPG FilNameExtension = iota
	JPEG
	PNG
	GIF
	TIFF
	BMP
)

var validateExtMap = map[string]FilNameExtension{
	".jpg":  JPG,
	".jpeg": JPEG,
	".png":  PNG,
	".gif":  GIF,
	".tiff": TIFF,
	".bmp":  BMP,
}

// validatorOperator はIValidatorOperatorの実装である。
type validatorOperator struct {
	operator ivalidator.ValidatorOperator
}

// NewValidatorOperator はvalidatorOperatorのインスタンスを生成する。
func NewValidatorOperator() *validatorOperator {
	return &validatorOperator{}
}

// ValidatoExt はIValidatorOperator.ValidateExtの実装である。
func (v *validatorOperator) ValidateExt(
	ext string,
) error {
	strings.ToLower(ext)

	_, ok := validateExtMap[ext]
	if !ok {
		return fmt.Errorf("invalid extension")
	}

	return nil
}
