package server

import "github.com/go-playground/validator"

type Login struct {
	Username string `validate:"required"`
	Knock    []int  `validate:"required"`
}

func (l *Login) Validate() error {
	return validator.New().Struct(l)
}
