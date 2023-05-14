package auth

import (
	"context"
	"encoding/json"
	"fmt"

	"go-zero-tutorial/service/cart/rpc/cartRpc/proto"
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
	cartResp, err := l.svcCtx.CartRpcClient.SearchUserCartTotal(l.ctx, &proto.SearchUserCartTotalRequest{UserId: uint64(uid)})
	jResp, _ := json.MarshalIndent(cartResp, "", "  ")
	logx.Infof("原始 cart rpc 返回结果 json 格式为： %v \n", string(jResp))
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
