package auth

import (
	"net/http"

	"go-zero-tutorial/service/userCenter/api/internal/logic/auth"
	"go-zero-tutorial/service/userCenter/api/internal/svc"
	"go-zero-tutorial/service/userCenter/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// curl --location --request POST '0.0.0.0:8801/v1/auth/login' \
// --header 'Content-Type: application/x-www-form-urlencoded' \
// --data-urlencode 'account=123' \
// --data-urlencode 'password=456'
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
