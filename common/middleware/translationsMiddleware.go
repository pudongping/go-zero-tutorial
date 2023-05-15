package middleware

import (
	"context"
	"net/http"

	"go-zero-tutorial/pkg/translator"
)

type TranslationsMiddleware struct {
}

func NewTranslationsMiddleware() *TranslationsMiddleware {
	return &TranslationsMiddleware{}
}

func (m *TranslationsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locale := r.Header.Get("locale")
		trans := translator.NewTranslator(locale)

		// 将 Translator 存储到全局上下文中，便于后续翻译时的使用
		ctx := r.Context()
		ctx = context.WithValue(ctx, translator.ValidatorTranslatorKey, trans)

		next(w, r.WithContext(ctx))
	}
}
