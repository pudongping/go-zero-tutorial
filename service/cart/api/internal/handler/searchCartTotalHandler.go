package handler

import (
	"net/http"

	"go-zero-tutorial/service/cart/api/internal/logic"
	"go-zero-tutorial/service/cart/api/internal/svc"
	"go-zero-tutorial/service/cart/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// curl --location --request GET '0.0.0.0:8802/carts/123/total' \
// --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzI0MjU1MTksImlhdCI6MTY3MjQxODMxOSwidXNlcklkIjoxfQ.oYS4x365X8hq1jCur2-CXusM9GxziOf3x9rKq7p7IMg' \
// --form 'account="alex"'
func searchCartTotalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchCartTotalRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSearchCartTotalLogic(r.Context(), svcCtx)
		resp, err := l.SearchCartTotal(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
