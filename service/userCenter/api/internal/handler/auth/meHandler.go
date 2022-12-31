package auth

import (
	"net/http"

	"go-zero-tutorial/service/userCenter/api/internal/logic/auth"
	"go-zero-tutorial/service/userCenter/api/internal/svc"
	"go-zero-tutorial/service/userCenter/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// curl --location --request GET '0.0.0.0:8801/v1/auth/me' \
// --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzI0ODk4NjMsImlhdCI6MTY3MjQ4MjY2MywidXNlcklkIjoxfQ.TsfbgGMOy5G-qgFHBMkAc0T5qGwzsgjIsWYJsGobexo'
func MeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthMeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewMeLogic(r.Context(), svcCtx)
		resp, err := l.Me(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
