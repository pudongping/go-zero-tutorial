package main

import (
	"flag"
	"fmt"
	"net/http"

	"go-zero-tutorial/common/httpRest"
	commonMiddleware "go-zero-tutorial/common/middleware"
	"go-zero-tutorial/service/userCenter/api/internal/config"
	"go-zero-tutorial/service/userCenter/api/internal/handler"
	"go-zero-tutorial/service/userCenter/api/internal/middleware"
	"go-zero-tutorial/service/userCenter/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// yaml 配置文件中使用环境变量时，则需要在使用 `conf.UseEnv()`
	// yaml 中使用环境变量示例：Name: ${ProjectName}
	// conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCors(), // 设置跨域
		// rest.WithCors("https://www.baidu.com"), // 设置单个域名跨域
		rest.WithNotFoundHandler(httpRest.NotFoundHandler()),        // 处理 404 情况
		rest.WithNotAllowedHandler(httpRest.NotAllowedHandler()),    // 处理请求方法不对时的调用
		rest.WithUnauthorizedCallback(httpRest.UnauthorizedHandler), // 接口没有认证成功时
	)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	// 全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Info("start 全局中间件")
			next(w, r)
			logx.Info("end 全局中间件")
		}
	})
	server.Use(new(middleware.CommonHeader).AddHeader())
	server.Use(commonMiddleware.NewTranslationsMiddleware().Handle) // 全局翻译中间件

	handler.RegisterHandlers(server, ctx)

	// logx.DisableStat()  // 关闭输出的统计 stat 日志

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
