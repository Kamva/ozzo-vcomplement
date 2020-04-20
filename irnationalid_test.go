package vcomplement

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIRNationalID(t *testing.T) {
	id := "3570093492"

	err := validation.Validate(&id, IRNationalID)

	assert.Nil(t, err)
}

func TestIRNationalIDInvalid(t *testing.T) {
	id := "3570093493"

	err := validation.Validate(&id, IRNationalID)

	assert.Equal(t,ErrIRNationalIDInvalid,err)
}

func TestIRNationalIDInvalidWithAlpha(t *testing.T) {
	id := "a570093492"

	err := validation.Validate(&id, IRNationalID)

	assert.Equal(t,ErrIRNationalIDInvalid,err)
}

