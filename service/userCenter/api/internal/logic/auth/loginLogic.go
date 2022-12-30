package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"go-zero-tutorial/service/userCenter/api/internal/svc"
	"go-zero-tutorial/service/userCenter/api/internal/types"
	"go-zero-tutorial/service/userCenter/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	if len(strings.TrimSpace(req.Account)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数不能为空")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Account)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户名不存在")
	default:
		return nil, err
	}

	if userInfo.Password != req.Password {
		return nil, errors.New("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.generateJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.UserId)
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResponse{
		LoginUserInfo: types.LoginUserInfo{
			UserId:    int(userInfo.UserId),
			Email:     userInfo.Email,
			Name:      userInfo.Name,
			Password:  "",
			Gender:    int8(userInfo.Gender),
			CreatedAt: int(userInfo.CreatedAt),
			UpdatedAt: int(userInfo.UpdatedAt),
		},
		AccessToken:  jwtToken,
		AccessExpire: accessExpire,
		RefreshTime:  now + accessExpire/2,
	}

	return
}

// 生成 token
func (l *LoginLogic) generateJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds            // 签名过期时间
	claims["iat"] = iat                      // 生成签名的时间
	claims["userId"] = userId                // 当前用户 id
	token := jwt.New(jwt.SigningMethodHS256) // 使用 HS256 算法生成的 token
	token.Claims = claims
	return token.SignedString([]byte(secretKey)) // 生成签名字符串
}
