package validation

import "gopkg.in/go-playground/validator.v9"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// It takes a struct as an argument, validates it, and returns an array of errors
func ValidateStruct(i interface{}) []*ErrorResponse {
	validate := validator.New()

	var errors []*ErrorResponse
	err := validate.Struct(i)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
