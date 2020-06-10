package vcomplement

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoney(t *testing.T) {
	values := []struct {
		Tag     string
		Value   string
		IsValid bool
	}{
		{"t1", "21000", true},
		{"t2", "21000.0", true},
		{"t3", "21000.384", true},
		{"t4", "0.0", true},
		{"t5", "0.324", true},
		{"t6", ".343", true},
		{"t7", "-.343", false},
		{"t8", "-1.343", false},
		{"t9", "-0.343", false},
	}

	for _, val := range values {
		err := validation.Validate(&val.Value, Money)
		if val.IsValid {
			assert.Nil(t, err, val.Tag)
		} else {
			assert.Error(t, err, val.Tag)
		}
	}

}
