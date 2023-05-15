package translator

import (
	"context"
	"errors"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_HK"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

const ValidatorTranslatorKey = "validatorTranslator"

func NewTranslator(locale string) ut.Translator {
	if "" == locale {
		locale = "zh"
	}

	uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New(), zh_Hant_HK.New())
	trans, _ := uni.GetTranslator(locale)

	validate := validator.New()
	switch locale {
	case "zh":
		// 调用 RegisterDefaultTranslations 方法将验证器和对应语言类型的 Translator 注册进来，实现验证器的多语言支持
		_ = zh_translations.RegisterDefaultTranslations(validate, trans)
		break
	case "en":
		_ = en_translations.RegisterDefaultTranslations(validate, trans)
		break
	default:
		_ = zh_translations.RegisterDefaultTranslations(validate, trans)
		break
	}

	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	return trans
}

func TranslateErrs(ctx context.Context, err error) error {
	if err == nil {
		return err
	}
	trans, _ := ctx.Value(ValidatorTranslatorKey).(ut.Translator)
	validateErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	var respErr string
	for k, v := range validateErrs.Translate(trans) {
		respErr += "field: " + k + " reason: " + v + " | "
	}
	return errors.New(respErr)
}
