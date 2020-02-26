package vcomplement

import (
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
	ErrInternalValidation = kitty.NewError(http.StatusInternalServerError, "OZVC_1_E_00",
		kitty.ReplyErrKeyInternalError, "internal error")

	ErrValidationError = kitty.NewError(http.StatusBadRequest, "OZVC_1_E_01",
		"invalid_input_data", "invalid input data")
)
