package ivalidator

// ValidatorOperator はvalidator.goを操作するインターフェース。
type ValidatorOperator interface {
	// ValidateExt は画像の拡張子のバリデーションチェックを行う。
	ValidateExt(ext string) error
}