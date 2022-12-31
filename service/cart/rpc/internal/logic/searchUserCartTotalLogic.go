package logic

import (
	"context"

	"go-zero-tutorial/service/cart/rpc/internal/svc"
	"go-zero-tutorial/service/cart/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserCartTotalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserCartTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserCartTotalLogic {
	return &SearchUserCartTotalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserCartTotalLogic) SearchUserCartTotal(in *proto.SearchUserCartTotalRequest) (*proto.SearchUserCartTotalResponse, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(in.UserId))
	if err != nil {
		return nil, err
	}

	return &proto.SearchUserCartTotalResponse{
		UserId:  uint64(user.UserId),
		Email:   user.Email,
		Account: user.Name,
	}, nil
}
