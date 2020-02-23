package vcomplement

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfirm(t *testing.T) {
	expected := "abc"
	val := "abc"

	err := validation.Validate(&val, Confirm(expected))

	assert.Nil(t, err)
}

func TestConfirmRealValuePointer(t *testing.T) {
	expected := "abc"
	val := "abc"

	err := validation.Validate(&val, Confirm(&expected))

	assert.Nil(t, err)
}

func TestConfirmError(t *testing.T) {
	expected := "123"
	val := "abc"

	err := validation.Validate(&val, Confirm(expected))

	assert.Equal(t, ErrInvalidConfirmation, err)
}
