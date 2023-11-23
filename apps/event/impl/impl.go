package impl

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/infraboard/maudit/apps/event"
	"github.com/infraboard/mcube/ioc/config/kafka"
	ioc_mongo "github.com/infraboard/mcube/ioc/config/mongo"
	go_kafka "github.com/segmentio/kafka-go"
)

var (
	// Service 服务实例
	svr = &impl{
		GroupId: "maudit.event.consumer",
		Topics:  []string{"maudit.event"},
	}
)

type impl struct {
	col *mongo.Collection
	log *zerolog.Logger
	kr  *go_kafka.Reader

	event.UnimplementedServiceServer
	ioc.ObjectImpl

	GroupId string   `toml:"group_id" json:"group_id" yaml:"group_id" env:"EVNET_GROUP_ID"`
	Topics  []string `toml:"topics" json:"topics" yaml:"topics" env:"EVNET_TOPICS" envSeparator:","`
}

func (i *impl) Init() error {
	i.log = logger.Sub(i.Name())

	i.col = ioc_mongo.DB().Collection(i.Name())
	indexs := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "save_at", Value: -1}},
		},
	}
	_, err := i.col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	i.log.Debug().Msgf("group_id: %s, topics: %s", i.GroupId, i.Topics)
	i.kr = kafka.ConsumerGroup(i.GroupId, i.Topics)
	go i.ConsumerEvent()
	return nil
}

// 对象的销毁
func (s *impl) Close(ctx context.Context) error {
	return s.kr.Close()
}

func (s *impl) Name() string {
	return event.AppName
}

func (s *impl) Registry(server *grpc.Server) {
	event.RegisterServiceServer(server, svr)
}

func init() {
	ioc.Controller().Registry(svr)
}
