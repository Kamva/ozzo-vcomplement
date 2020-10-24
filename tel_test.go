package vcomplement

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIRTel(t *testing.T) {
	phone := "07735451354"

	err := validation.Validate(&phone, IRTel)
	assert.Nil(t, err)
}

func TestIRTel0(t *testing.T) {
	phone := "07735451354"

	err := validation.Validate(&phone, IRTel0)
	assert.Nil(t, err)
}

func TestIRTel98(t *testing.T) {
	phone := "+987735451354"

	err := validation.Validate(&phone, IRTel98)
	assert.Nil(t, err)
}

func TestInvalidIRTelError(t *testing.T) {
	phone := "077-35451354" // with dash between digits

	err := validation.Validate(&phone, IRTel98)

	assert.Equal(t, ErrIRTelInvalid, err)
}
