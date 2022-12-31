package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-tutorial/service/cart/api/internal/logic"
	"go-zero-tutorial/service/cart/api/internal/svc"
)

func cartTotalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCartTotalLogic(r.Context(), svcCtx)
		err := l.CartTotal()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
