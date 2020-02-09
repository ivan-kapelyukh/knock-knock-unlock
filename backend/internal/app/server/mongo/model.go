package mongo

import "github.com/go-playground/validator"

type User struct {
	Username string  `validate:"required"`
	Knocks   [][]int `validate:"minlength"`
}

const minimumKnocks = 3

func (u *User) Validate() error {
	validate := validator.New()

	err := validate.RegisterValidation("minlength", atLeastMinimumLength)

	if err != nil {
		return err
	}

	return validate.Struct(u)
}

func atLeastMinimumLength(fl validator.FieldLevel) bool {
	return fl.Field().Len() >= minimumKnocks
}
