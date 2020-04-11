package vcomplement

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSocialNetwork(t *testing.T) {
	social := "my-account"

	err := validation.Validate(&social, SocialNetworkIdentifier)

	assert.Nil(t, err)
}

