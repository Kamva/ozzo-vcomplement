package vcomplement

import (
	"github.com/BurntSushi/toml"
	"github.com/Kamva/gutil"
	"github.com/Kamva/kitty"
	"github.com/Kamva/kitty/kittytranslator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

type ABC struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func newTranslator() kitty.Translator {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	gutil.Must(bundle.LoadMessageFile(gutil.SourcePath() + "/testdata/en.toml"))
	return kittytranslator.NewI18nDriver(bundle, i18n.NewLocalizer(bundle, "en"), []string{})
}

func TestSingleFieldValidationInvalid(t *testing.T) {
	vt := NewKittyDriverErrorTranslator(newTranslator())
	bag := &TranslateBag{singleMessage: "abc_alpha"}

	name := "123"
	res, err := vt.Translate(validation.Validate(&name, is.Alpha))

	assert.Nil(t, err)
	assert.Equal(t, bag, res)
}

func TestSingleFieldValidationValid(t *testing.T) {
	vt := NewKittyDriverErrorTranslator(newTranslator())

	name := "abc"
	res, err := vt.Translate(validation.Validate(&name, is.Alpha))

	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestSingleFieldValidationDefaultMessage(t *testing.T) {
	vt := NewKittyDriverErrorTranslator(newTranslator())
	bag := &TranslateBag{singleMessage: is.ErrURL.Message()}

	name := "abc"
	res, err := vt.Translate(validation.Validate(&name, is.URL))

	assert.Nil(t, err)
	assert.Equal(t, bag, res)
}

func TestSingleFieldValidationInvalidErr(t *testing.T) {
	vt := NewKittyDriverErrorTranslator(newTranslator())

	name := "abc"
	res, err := vt.Translate(validation.Validate(&name, validation.Min(10)))

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestStructValidation(t *testing.T) {
	vt := NewKittyDriverErrorTranslator(newTranslator())

	tests := []struct {
		tag  string
		data ABC
		bag  *TranslateBag
	}{
		{
			"t1", ABC{Name: "123", Age: 4},
			&TranslateBag{groupMessages: map[string]*TranslateBag{
				"name": {singleMessage: "abc_alpha"},
				"age":  {singleMessage: "abc_min"},
			}},
		},
		{
			"t2", ABC{Name: "123", Age: 11},
			&TranslateBag{groupMessages: map[string]*TranslateBag{
				"name": {singleMessage: "abc_alpha"},
			}},
		},
		{"t3", ABC{Name: "abc", Age: 11}, nil},
	}

	for _, test := range tests {
		d := &test.data

		res, err := vt.Translate(validation.ValidateStruct(d,
			validation.Field(&d.Name, is.Alpha),
			validation.Field(&d.Age, validation.Min(10)),
		))

		assert.Nil(t, err, test.tag)
		assert.Equal(t, test.bag, res, test.tag)
	}

}
func TestStructValidationDefaultMessage(t *testing.T) {
	vt := NewKittyDriverErrorTranslator(newTranslator())

	tests := []struct {
		tag  string
		data ABC
		bag  *TranslateBag
	}{
		{
			"t1", ABC{Name: "123", Age: 11},
			&TranslateBag{groupMessages: map[string]*TranslateBag{
				"name": {singleMessage: is.ErrURL.Error()},
				"age": {singleMessage: validation.ErrMaxLessEqualThanRequired.SetParams(map[string]interface{}{
					"threshold": 10,
				}).Error()},
			}},
		},

		{
			"t2", ABC{Name: "abc", Age: 4},
			&TranslateBag{groupMessages: map[string]*TranslateBag{
				"name": {singleMessage: is.ErrURL.Error()},
			}},
		},

		{"t3", ABC{Name: "http://abc.com", Age: 4}, nil},
	}

	for _, test := range tests {
		d := &test.data

		res, err := vt.Translate(validation.ValidateStruct(d,
			validation.Field(&d.Name, is.URL),
			validation.Field(&d.Age, validation.Max(10)),
		))

		assert.Nil(t, err, test.tag)
		assert.Equal(t, test.bag, res, test.tag)
	}
}

func TestStructValidationInvalidError(t *testing.T) {
	vt := NewKittyDriverErrorTranslator(newTranslator())

	tests := []struct {
		tag  string
		data ABC
	}{
		{"t1", ABC{Name: "123", Age: 11}},
		{"t2", ABC{Name: "abc", Age: 4}},
		{"t3", ABC{Name: "http://abc.com", Age: 4}},
	}

	for _, test := range tests {
		d := &test.data

		res, err := vt.Translate(validation.ValidateStruct(d,
			validation.Field(&d.Name, is.URL),
			validation.Field(&d.Age, is.JSON),
		))

		assert.NotNil(t, err, test.tag)
		assert.Nil(t, res, test.tag)
	}
}
