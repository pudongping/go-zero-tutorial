package logic

import (
	"context"
	"fmt"

	"go-zero-tutorial/service/cart/rpc/cartRpc/internal/svc"
	"go-zero-tutorial/service/cart/rpc/cartRpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
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

	// 尝试从 metadata 中获取 rpc 客户端携带过来的参数
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		tmp := md.Get("uid")
		fmt.Printf("metadata tmp : %+v \n", tmp)
	}

	fmt.Println("SearchUserCartTotalLogic SearchUserCartTotal Fun")
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
