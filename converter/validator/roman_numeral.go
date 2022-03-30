package validator

import (
	"converter/romantoarabic"
	"github.com/go-playground/validator/v10"
)

var IsRomanNumeral validator.Func = func(fl validator.FieldLevel) bool {
	return romantoarabic.IsRomanNumber(fl.Field().String())
}
