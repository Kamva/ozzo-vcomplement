package vcomplement

import (
	"errors"
	"github.com/kamva/hexa"
	"net/http"
)

//--------------------------------
// Translation Errors
//--------------------------------

// Error code description:
// OZVC = OZZO validation Complement template (package or project name)
// tr = errors about translation section (identify some part in application)
// E = Error (type of code : error|response|...)
// 00 = error number zero (id of code in that part and type)

var (
	ErrInternalValidation = hexa.NewError(
		http.StatusInternalServerError,
		"lib.translation.internal_error",
		errors.New("internal error"),
	)

	ErrValidationError = hexa.NewError(
		http.StatusBadRequest,
		"lib.translation.invalid_input_data_error",
		errors.New("invalid input data"),
	)
)
