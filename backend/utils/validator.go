package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateStruct(s interface{}) error {

	validate = validator.New()

	err := validate.Struct(s)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
