package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/maudit/apps/event"
)

var (
	h = &handler{}
)

type handler struct {
	service event.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(h.Name())
	h.service = app.GetGrpcApp(event.AppName).(event.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return event.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	rr := r.ResourceRouter("events")

	rr.BasePath("events")
}

func init() {
	app.RegistryHttpApp(h)
}
