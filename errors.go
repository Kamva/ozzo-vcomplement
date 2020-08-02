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
	ErrInternalValidation = hexa.NewError(http.StatusInternalServerError, "ozvc.tr.e.0",
		hexa.ErrKeyInternalError, errors.New("internal error"))

	ErrValidationError = hexa.NewError(http.StatusBadRequest, "ozvc.tr.e.1",
		"err_invalid_input_data", errors.New("invalid input data"))
)
