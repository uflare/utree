package main

import "github.com/go-playground/validator"

// Validator - the request validator
type Validator struct {
	validator *validator.Validate
}

// Validate - validate the specified input
func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
