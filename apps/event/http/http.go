package http

import (
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"

	"github.com/infraboard/maudit/apps/event"
)

var (
	h = &handler{}
)

type handler struct {
	service event.ServiceServer
	log     *zerolog.Logger

	ioc.ObjectImpl
}

func (h *handler) Config() error {
	h.log = logger.Sub(h.Name())
	h.service = ioc.Controller().Get(event.AppName).(event.ServiceServer)
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
	ioc.Api().Registry(h)
}
