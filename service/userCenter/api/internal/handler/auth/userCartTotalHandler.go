package auth

import (
	"net/http"

	"go-zero-tutorial/service/userCenter/api/internal/logic/auth"
	"go-zero-tutorial/service/userCenter/api/internal/svc"
	"go-zero-tutorial/service/userCenter/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// curl --location --request GET '0.0.0.0:8801/v1/auth/userCartTotal' \
// --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUzNDM0OTcsImlhdCI6MTY4NDA0NzQ5NywidXNlcklkIjoxfQ.1owRkFecBoxtYxEZbSPx1LCe7lYesDtNS6gErLg44h0'
func UserCartTotalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCartTotalRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewUserCartTotalLogic(r.Context(), svcCtx)
		resp, err := l.UserCartTotal(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
