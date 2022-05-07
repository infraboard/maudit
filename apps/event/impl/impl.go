package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/maudit/apps/event"
	"github.com/infraboard/maudit/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger

	event.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	col := db.Collection(s.Name())
	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "save_at", Value: bsonx.Int32(-1)}},
		},
	}
	_, err = col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}
	s.col = col
	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *service) Name() string {
	return event.AppName
}

func (s *service) Registry(server *grpc.Server) {
	event.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
