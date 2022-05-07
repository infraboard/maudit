package nats_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/maudit/apps/event"
	"github.com/infraboard/maudit/apps/event/bus/broker/nats"
)

func TestPubSub(t *testing.T) {
	should := assert.New(t)
	log := zap.L().Named("Nats Bus")

	oe := &event.OperateEventData{
		Session: "xxx1",
		Account: "test",
	}
	sourceEvent, err := event.NewProtoOperateEvent(oe)
	should.NoError(err)

	nc := nats.NewDefaultConfig()
	b, err := nats.NewBroker(nc)
	b.Debug(log)
	should.NoError(err)

	should.NoError(b.Connect())

	// 后台先起进程来订阅处理
	err = b.Sub("test", func(topic string, e *event.Event) error {
		should.Equal(sourceEvent.Id, e.Id)
		target := &event.OperateEventData{}
		err := e.ParseBoby(target)
		should.NoError(err)
		should.Equal(oe.Account, target.Account)
		log.Info(target)
		return nil
	})
	should.NoError(err)

	should.NoError(b.Pub("test", sourceEvent))
	should.NoError(b.Disconnect())
}

func init() {
	zap.DevelopmentSetup()
}
