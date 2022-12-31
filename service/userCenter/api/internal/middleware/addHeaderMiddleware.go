package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/rest"
)

type CommonHeader struct {
}

func (ch *CommonHeader) GenRequestId() string {
	return stringx.Randn(24)
}

func (ch *CommonHeader) AddHeader() rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Info("start CommonHeader Middleware")
			w.Header().Add("X-Request-ID", ch.GenRequestId())
			w.Header().Add("X-Server", "go-zero")
			next(w, r)
			logx.Info("end CommonHeader Middleware")
		}
	}
}
