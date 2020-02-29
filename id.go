package vcomplement

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

//--------------------------------
// Iran Phone number.
//--------------------------------

// ErrInvalidID is the default invalid id validation error
var ErrInvalidID = validation.NewError("validation_invalid_id", "ID value is invalid")

// ID function return new validator to validate id.
func ID(validationErr error) validation.Rule {
	return validation.By(func(value interface{}) error {
		value, isNil := validation.Indirect(value)
		if isNil || validation.IsEmpty(value) {
			return nil
		}

		if validationErr != nil {
			return ErrInvalidID.SetMessage(validationErr.Error())
		}

		return nil
	})
}
