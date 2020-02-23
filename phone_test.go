package vcomplement

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIRPhpne(t *testing.T) {
	phone := "09361234567"

	err := validation.Validate(&phone, IRPhone)

	assert.Nil(t, err)
}

func TestInvalidIRPhoneError(t *testing.T) {
	phone := "9361234567"

	err := validation.Validate(&phone,IRPhone)

	assert.Equal(t, ErrIRPhoneInvalid, err)
}
