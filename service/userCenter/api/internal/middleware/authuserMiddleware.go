package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthUserMiddleware struct {
}

func NewAuthUserMiddleware() *AuthUserMiddleware {
	return &AuthUserMiddleware{}
}

func (m *AuthUserMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("开始进入 AuthUserMiddleware 中间件")
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("start AuthUserMiddleware")
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
		logx.Info("end AuthUserMiddleware")
	}
}
