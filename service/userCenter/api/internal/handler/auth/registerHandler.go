package auth

import (
	"net/http"

	"go-zero-tutorial/pkg/translator"
	"go-zero-tutorial/service/userCenter/api/internal/logic/auth"
	"go-zero-tutorial/service/userCenter/api/internal/svc"
	"go-zero-tutorial/service/userCenter/api/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 使用第三方验证器：github.com/go-playground/validator/v10
		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			// 添加翻译器开始
			err = translator.TranslateErrs(r.Context(), err)
			// 添加翻译器结束
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
