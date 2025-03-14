package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/setting"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	notify.UnimplementedRPCServer

	user    user.Service
	log     logger.Logger
	setting setting.Service
	domain  domain.Service
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.col = db.Collection(s.Name())
	s.log = zap.L().Named(s.Name())
	s.setting = app.GetInternalApp(setting.AppName).(setting.Service)
	s.user = app.GetInternalApp(user.AppName).(user.Service)
	s.domain = app.GetInternalApp(domain.AppName).(domain.Service)
	return nil
}

func (s *service) Name() string {
	return notify.AppName
}

func (s *service) Registry(server *grpc.Server) {
	notify.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
