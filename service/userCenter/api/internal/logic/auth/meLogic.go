package auth

import (
	"context"
	"encoding/json"
	"errors"

	"go-zero-tutorial/service/userCenter/api/internal/svc"
	"go-zero-tutorial/service/userCenter/api/internal/types"
	"go-zero-tutorial/service/userCenter/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeLogic {
	return &MeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeLogic) Me(req *types.AuthMeRequest) (resp *types.AuthMeResponse, err error) {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, errors.New("解析失败：" + err.Error())
	}
	if userId == 0 {
		return nil, errors.New("无法获取当前用户 id")
	}
	logx.Infof("从 header 头中解析的 token 之后的 userId: %v", userId)

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("找不到当前用户信息")
		}
		return nil, err
	}

	resp = &types.AuthMeResponse{
		UserInfo: types.UserInfo{
			UserId:    int(userInfo.UserId),
			Email:     userInfo.Email,
			Name:      userInfo.Name,
			Password:  userInfo.Password,
			Gender:    int8(userInfo.Gender),
			CreatedAt: int(userInfo.CreatedAt),
			UpdatedAt: int(userInfo.UpdatedAt),
		},
	}

	return
}
