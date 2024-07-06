package providerwire

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslate "github.com/go-playground/validator/v10/translations/en"
)

func provideValidator(t ut.Translator) *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	_ = entranslate.RegisterDefaultTranslations(v, t)

	return v
}

func provideTranslator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	return trans
}
