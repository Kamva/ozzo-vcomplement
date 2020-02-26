package vcomplement

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

//--------------------------------
// Iran Phone number.
//--------------------------------

// ErrIRPhoneInvalid is the default IRPhone validation rules error.
var ErrIRPhoneInvalid = validation.NewError("validation_ir_phone_invalid", "Iran phone number is invalid")

// IRPhone is the iran phone number validation rule.
var IRPhone = validation.Match(regexp.MustCompile("^(\\+98|0)9\\d{9}$")).ErrorObject(ErrIRPhoneInvalid)
var IRPhone98 = validation.Match(regexp.MustCompile("^\\+989\\d{9}$")).ErrorObject(ErrIRPhoneInvalid)
var IRPhone0 = validation.Match(regexp.MustCompile("^09\\d{9}$")).ErrorObject(ErrIRPhoneInvalid)
