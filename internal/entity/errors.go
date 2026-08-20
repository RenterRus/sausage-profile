package entity

import "errors"

var (
	ErrParametrNoFound = errors.New("params not found")
	ErrCodeInvalid     = errors.New("code invalid")
)
