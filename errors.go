package vcomplement

import (
	"errors"
	"github.com/Kamva/kitty"
	"net/http"
)

//--------------------------------
// Translation Errors
//--------------------------------

// Error code description:
// OZVC = OZZO validation Complement template (package or project name)
// 1 = errors about translation section (identify some part in application)
// E = Error (type of code : error|response|...)
// 00 = error number zero (id of code in that part and type)

var (
	ErrInternalValidation = kitty.NewError(http.StatusInternalServerError, "ozvc.tr.e.0",
		kitty.ErrKeyInternalError, errors.New("internal error"))

	ErrValidationError = kitty.NewError(http.StatusBadRequest, "ozvc.tr.e.1",
		"err_invalid_input_data", errors.New("invalid input data"))
)
