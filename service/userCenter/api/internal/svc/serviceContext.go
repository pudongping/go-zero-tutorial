package svc

import (
	"context"
	"fmt"

	"go-zero-tutorial/service/cart/rpc/cartRpc/cart"
	"go-zero-tutorial/service/userCenter/api/internal/config"
	"go-zero-tutorial/service/userCenter/api/internal/middleware"
	"go-zero-tutorial/service/userCenter/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserModel          model.UserModel
	AuthUserMiddleware rest.Middleware
	CartRpcClient      cart.Cart // 如果是一个服务一个项目的话，需要重新生成 proto 文件，这里为了演示方便，直接使用 cart 服务下的 grpc 客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})

	// cart rpc 连接方式
	cartRpcClient1 := zrpc.MustNewClient(c.CartRpcConf, zrpc.WithUnaryClientInterceptor(TestClientInterceptor)) // 1、etcd 服务发现 2、endpoints 直连方式
	// cartRpcClient2, _ := zrpc.NewClientWithTarget("0.0.0.0:8901") // 3、ip 直连模式

	return &ServiceContext{
		Config:             c,
		RedisClient:        redisClient,
		UserModel:          model.NewUserModel(mysqlConn, c.CacheRedis),
		AuthUserMiddleware: middleware.NewAuthUserMiddleware().Handle,
		CartRpcClient:      cart.NewCart(cartRpcClient1),
	}
}

func TestClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	// 尝试将值通过 metadata 传递给 rpc 服务端
	md := metadata.New(map[string]string{
		"uid": "1688",
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	fmt.Printf("grpc 客户端拦截器 start ====> \n")
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("grpc 客户端拦截器 end ====> \n")

	return err
}
