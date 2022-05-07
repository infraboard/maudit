package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/maudit/apps/event"
)

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	req := &event.SaveEventRequest{}
	ins, err := h.service.SaveEvent(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
	return
}

func (h *handler) QueryEvent(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := event.NewQueryEventkRequest(page)

	ins, err := h.service.QueryEvent(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
	return
}
