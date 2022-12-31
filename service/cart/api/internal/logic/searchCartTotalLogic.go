package logic

import (
	"context"

	"go-zero-tutorial/service/cart/api/internal/svc"
	"go-zero-tutorial/service/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCartTotalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchCartTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCartTotalLogic {
	return &SearchCartTotalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCartTotalLogic) SearchCartTotal(req *types.SearchCartTotalRequest) (resp *types.SearchCartTotalResponse, err error) {
	// 拿到 jwt 中的载荷
	payload := map[string]interface{}{
		"userId": l.ctx.Value("userId"),
	}
	// {"@timestamp":"2022-12-31T01:01:49.292+08:00","caller":"logic/searchCartTotalLogic.go:31","content":"parse jwt payload: map[userId:1]","level":"info"}
	logx.Infof("parse jwt payload: %v", payload)

	return
}
