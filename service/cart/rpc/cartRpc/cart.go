package main

import (
	"context"
	"flag"
	"fmt"

	"go-zero-tutorial/service/cart/rpc/cartRpc/internal/config"
	"go-zero-tutorial/service/cart/rpc/cartRpc/internal/server"
	"go-zero-tutorial/service/cart/rpc/cartRpc/internal/svc"
	"go-zero-tutorial/service/cart/rpc/cartRpc/proto"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/cart.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		proto.RegisterCartServer(grpcServer, server.NewCartServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// grpc 服务端拦截器
	s.AddUnaryInterceptors(TestServerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func TestServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("grpc 服务端拦截器 start ====> \n")
	fmt.Printf("req ====> %+v \n", req)
	fmt.Printf("info ====> %+v \n", info)
	// req ====> user_id:1
	// info ====> &{Server:0xc0004975e0 FullMethod:/proto.Cart/SearchUserCartTotal}

	resp, err = handler(ctx, req)

	fmt.Printf("grpc 服务端拦截器 end ====> \n")

	return
}
