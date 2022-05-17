package impl_test

import (
	"testing"

	"github.com/infraboard/maudit/apps/event"
	"github.com/infraboard/maudit/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	svc event.ServiceServer
)

func TestSaveEvent(t *testing.T) {

}

func init() {
	zap.DevelopmentSetup()
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
	svc = app.GetGrpcApp(event.AppName).(event.ServiceServer)
}
