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
	accountNumber1 := "73483"
	accountNumber2 := "8374837843827383"

	assert.Nil(t, validation.Validate(&accountNumber1, IRAccountNumber))
	assert.Nil(t, validation.Validate(&accountNumber2, IRAccountNumber))
}

func TestInvalidIRPAccountNumberError(t *testing.T) {
	accountNumber := "9361234567"

	err := validation.Validate(&accountNumber, IRAccountNumber)

	assert.Equal(t, ErrIRAccountNumberInvalid, err)
}
