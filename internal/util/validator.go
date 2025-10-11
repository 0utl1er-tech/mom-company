package util

import (
	"buf.build/go/protovalidate"
	"google.golang.org/protobuf/proto"
)

// Validator はprotovalidateのバリデーターを管理します
type Validator struct {
	validator protovalidate.Validator
}

// NewValidator は新しいValidatorを作成します
func NewValidator() (*Validator, error) {
	validator, err := protovalidate.New()
	if err != nil {
		return nil, err
	}
	return &Validator{validator: validator}, nil
}

// Validate はメッセージをバリデーションします
func (v *Validator) Validate(msg proto.Message) error {
	return v.validator.Validate(msg)
}

// GlobalValidator はグローバルなバリデーターインスタンスです
var GlobalValidator *Validator

// InitValidator はグローバルバリデーターを初期化します
func InitValidator() error {
	validator, err := NewValidator()
	if err != nil {
		return err
	}
	GlobalValidator = validator
	return nil
}

// ValidateMessage はメッセージをバリデーションします（グローバルバリデーターを使用）
func ValidateMessage(msg proto.Message) error {
	if GlobalValidator == nil {
		if err := InitValidator(); err != nil {
			return err
		}
	}
	return GlobalValidator.Validate(msg)
}
