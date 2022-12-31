package auth

import (
	"context"
	"encoding/json"
	"fmt"

	"go-zero-tutorial/service/cart/rpc/proto"
	"go-zero-tutorial/service/userCenter/api/internal/svc"
	"go-zero-tutorial/service/userCenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCartTotalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCartTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCartTotalLogic {
	return &UserCartTotalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCartTotalLogic) UserCartTotal(req *types.UserCartTotalRequest) (resp *types.UserCartTotalResponse, err error) {
	userId := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("当前的用户 id 为： %s", userId)
	uid, err := userId.Int64()
	if err != nil {
		return nil, err
	}

	// cart rpc
	cartResp, err := l.svcCtx.CartRpc.SearchUserCartTotal(l.ctx, &proto.SearchUserCartTotalRequest{UserId: uint64(uid)})
	logx.Infof("cart rpc 返回的结果为： %+v", cartResp)
	if err != nil {
		return nil, err
	}

	resp = &types.UserCartTotalResponse{
		UserId:    cartResp.UserId,
		Email:     cartResp.Email,
		Account:   cartResp.Account,
		CartTotal: cartResp.CartTotal,
	}

	return
}
