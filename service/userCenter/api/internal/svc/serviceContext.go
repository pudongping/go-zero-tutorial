package svc

import (
	"go-zero-tutorial/service/cart/rpc/cart"
	"go-zero-tutorial/service/userCenter/api/internal/config"
	"go-zero-tutorial/service/userCenter/api/internal/middleware"
	"go-zero-tutorial/service/userCenter/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	UserModel          model.UserModel
	AuthUserMiddleware rest.Middleware
	CartRpc            cart.Cart // 如果是一个服务一个项目的话，需要重新生成 proto 文件，这里为了方便，直接使用 cart 服务下的 grpc 客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	// cart rpc 连接方式
	cartRpcClient1 := zrpc.MustNewClient(c.CartRpc) // etcd 服务发现
	// cartRpcClient2, _ := zrpc.NewClientWithTarget("0.0.0.0:8901") // ip 直连模式

	return &ServiceContext{
		Config:             c,
		UserModel:          model.NewUserModel(mysqlConn, c.CacheRedis),
		AuthUserMiddleware: middleware.NewAuthUserMiddleware().Handle,
		CartRpc:            cart.NewCart(cartRpcClient1),
	}
}
