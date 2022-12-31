package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-tutorial/service/cart/api/internal/svc"
)

type CartTotalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartTotalLogic {
	return &CartTotalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartTotalLogic) CartTotal() error {
	// todo: add your logic here and delete this line

	return nil
}
