package vcomplement

import (
	"github.com/Kamva/gutil"
	"github.com/Kamva/kitty"
	"github.com/Kamva/tracer"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// Translator is the ozzo-validation error translator.
	Translator interface {
		// Translate translate validation error.
		Translate(err error) (*TranslateBag, error)

		// Wrap translated messages in a kitty Error, if err is nil, return nil.
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

// NewKittyDriverErrorTranslator returns new instance of kittyTranslator
//that translate ozzo-validation errors.
func NewKittyDriverErrorTranslator(t kitty.Translator) Translator {
	return &kittyTranslator{t: t}
}

func (t *kittyTranslator) translateErr(err validation.Error) (string, error) {
	val, e := t.t.TranslateDefault(err.Code(), err.Message(), gutil.MapToKeyValue(err.Params())...)
	return val, tracer.Trace(e)
}

func (t *kittyTranslator) Translate(err error) (*TranslateBag, error) {
	bag := new(TranslateBag)

	if e, ok := tracer.Cause(err).(validation.Error); ok {
		msg, err := t.translateErr(e)
		bag.SetSingleMsg(msg)
		return bag, tracer.Trace(err)
	}

	if es, ok := tracer.Cause(err).(validation.Errors); ok {
		for k, e := range es {
			errBag, err := t.Translate(e)

			if err != nil {
				return nil, tracer.Trace(err)
			}

			bag.AddMsgToGroup(k, errBag)
		}

		return bag, nil
	}

	// otherwise return just empty bag and error (if error is internal
	// error, so user can detect it, otherwise get empty bag that to
	// detect that validation does not have any error).
	return bag, err
}

// Wrap bag translated error messages to kitty Error.
// if bag be empty, returns nil.
func (t *kittyTranslator) WrapTranslationByError(err error) kitty.Error {
	bag, err := t.Translate(err)

	if err != nil {
		return ErrInternalValidation.SetError(tracer.Trace(err))
	}

	// Bag can be nil (in case of valid data)
	if bag.IsEmpty() {
		return nil
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

// IsEmpty specify that error bag is empty or not.
func (t *TranslateBag) IsEmpty() bool {
	return t.singleMessage == "" && len(t.groupMessages) == 0
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

	return messages
}

// TValidate get a translator and validatable interface, validate and return kitty error.
func TValidate(t Translator, v validation.Validatable) error {
	return t.WrapTranslationByError(v.Validate())
}

// TValidateBy validate by provided translator and check to detect right driver.
func TValidateByKitty(t kitty.Translator, v validation.Validatable) error {
	return TValidate(NewKittyDriverErrorTranslator(t.(kitty.Translator)), v)
}

// Assert kittyTranslator implements the Translator.
var _ Translator = &kittyTranslator{}
