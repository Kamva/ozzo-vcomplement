package vcomplement

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
)

func TestIRCardNumber(t *testing.T) {
	cardNumber := "5859831012190731"

	err := validation.Validate(&cardNumber, IRCardNumber)

	assert.Nil(t, err)
}

func TestInvalidIRPCardNumberError(t *testing.T) {
	cardNumber := "9361234567"

	err := validation.Validate(&cardNumber, IRCardNumber)

	assert.Equal(t, ErrIRCardNumberInvalid, err)
}

func TestIRAccountNumber(t *testing.T) {
	accountNumber1 := "73383"
	accountNumber2 := "8374837843827383"
	accountNumber3 := "8374837843827383123" // length 19

	assert.Nil(t, validation.Validate(&accountNumber1, IRAccountNumber))
	assert.Nil(t, validation.Validate(&accountNumber2, IRAccountNumber))
	assert.Nil(t, validation.Validate(&accountNumber3, IRAccountNumber))
}

func TestInvalidIRPAccountNumberError(t *testing.T) {
	length4 := "9312"
	length20 := "93129312931293129312"

	err := validation.Validate(&length4, IRAccountNumber)
	assert.Equal(t, ErrIRAccountNumberInvalid, err)

	err = validation.Validate(&length20, IRAccountNumber)
	assert.Equal(t, ErrIRAccountNumberInvalid, err)
}

func TestIRIBAN(t *testing.T) {
	iban := "IR830120000000000055771565"

	assert.Nil(t, validation.Validate(&iban, IRIBAN))
}

func TestInvalidIRIBANError(t *testing.T) {
	length25 := "IR83012000000000055771565"
	invalidChar := "IR83012000000000055771565d"
	length27 := "IR8301200000000005577153d53"
	withoutIR := "US83012000000000055771565"

	err := validation.Validate(&length25, IRIBAN)
	assert.Equal(t, ErrIRIBANNumberInvalid, err)

	err = validation.Validate(&invalidChar, IRIBAN)
	assert.Equal(t, ErrIRIBANNumberInvalid, err)

	err = validation.Validate(&length27, IRIBAN)
	assert.Equal(t, ErrIRIBANNumberInvalid, err)

	err = validation.Validate(&withoutIR, IRIBAN)
	assert.Equal(t, ErrIRIBANNumberInvalid, err)
}
