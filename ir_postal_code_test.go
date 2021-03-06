package vcomplement


import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIRPostalCode(t *testing.T) {
	p := "7741634183"

	err := validation.Validate(&p, IRPostalCode)

	assert.Nil(t, err)
}

func TestIRPostalCodeInvalid(t *testing.T) {
	p := "563798372"

	err := validation.Validate(&p, IRPostalCode)

	assert.Equal(t,ErrIRPostalCodeInvalid,err)
}

func TestIRPostalCodeInvalidWithAlpha(t *testing.T) {
	p := "abc4567890"

	err := validation.Validate(&p, IRPostalCode)

	assert.Equal(t,ErrIRPostalCodeInvalid,err)
}

