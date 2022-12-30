package svc

import (
	"go-zero-tutorial/service/userCenter/api/internal/config"
	"go-zero-tutorial/service/userCenter/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(mysqlConn, c.CacheRedis),
	}
}
