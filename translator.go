package vcomplement

import (
	"github.com/Kamva/gutil"
	"github.com/Kamva/kitty"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// Translator is the ozzo-validation error translator.
	Translator interface {
		Translate(err error) (*TranslateBag, error)
		WrapTranslationByError(err error) kitty.Error
	}

	// TranslateBag is the bag contains translated validation errors or error.
	TranslateBag struct {
		singleMessage string
		groupMessages map[string]*TranslateBag
	}
)

type kittyTranslator struct {
	t kitty.Translator
}

// NewKittyErrorTranslator returns new instance of kittyTranslator
//that translate ozzo-validation errors.
func NewKittyErrorTranslator(t kitty.Translator) Translator {
	return &kittyTranslator{t: t}
}

func (t *kittyTranslator) translateErr(err validation.Error) (string, error) {
	return t.t.TranslateDefault(err.Code(), err.Message(), gutil.MapToKeyValue(err.Params())...)
}

func (t *kittyTranslator) Translate(err error) (*TranslateBag, error) {
	bag := new(TranslateBag)

	if e, ok := err.(validation.Error); ok {
		msg, err := t.translateErr(e)
		bag.SetSingleMsg(msg)
		return bag, err
	}

	if es, ok := err.(validation.Errors); ok {
		for k, e := range es {
			errBag, err := t.Translate(e)

			if err != nil {
				return nil, err
			}

			bag.AddMsgToGroup(k, errBag)
		}
	}

	// otherwise return just simple error.
	return nil, err
}

func (t *kittyTranslator) WrapTranslationByError(err error) kitty.Error {
	bag, err := t.Translate(err)

	if err != nil {
		return ErrInternalValidation.SetError(err.Error())
	}

	return ErrValidationError.SetData(bag.Map(true).(map[string]interface{}))
}

// SetSingleMsg set single error message.
func (t *TranslateBag) SetSingleMsg(msg string) {
	t.singleMessage = msg
}

// AddMsgToGroup add message to the group
func (t *TranslateBag) AddMsgToGroup(key string, msg *TranslateBag) {
	if t.groupMessages == nil {
		t.groupMessages = map[string]*TranslateBag{}
	}

	t.groupMessages[key] = msg
}

// TOMap fucntion convert TranslateBag to a map[string]interface{}.
// but if bag just has a single message it check if forceMap is true,
// return map["error"]=<message>, otherwise returns just string message.
//
// possible values:
// - map[string]interface : when TranslateBag contains group of messages or
//	 contains single message with forceMap=true.
// - string: when TranslateBag contains singleMessage with forceMap=false.
// - nil : When TranslateBag does not hav single message nor group of messages.
func (t *TranslateBag) Map(forceMap bool) interface{} {
	if t.singleMessage != "" {
		if forceMap {
			return map[string]interface{}{"error": t.singleMessage}
		}

		return t.singleMessage
	}

	messages := map[string]interface{}{}

	for k, v := range t.groupMessages {
		messages[k] = v.Map(false)
	}

	return nil
}

// Assert kittyTranslator implements the Translator.
var _ Translator = &kittyTranslator{}
