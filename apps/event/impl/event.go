package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/maudit/apps/event"
	"github.com/infraboard/mcube/exception"
)

func (s *impl) SaveEvent(ctx context.Context, req *event.SaveEventRequest) (*event.SaveReponse, error) {
	ins, err := req.ParseEvent()
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", ins)

	if _, err := s.col.InsertMany(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted event(%s) document error, %s",
			req.Ids(), err)
	}

	resp := event.NewSaveReponse()
	resp.AddSuccess(req.Ids()...)
	return resp, nil
}

func (s *impl) QueryEvent(ctx context.Context, req *event.QueryEventRequest) (*event.OperateEventSet, error) {
	r := newQueryEventRequest(req)
	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find event error, error is %s", err)
	}

	set := event.NewOperateEventSet()
	// 循环
	for resp.Next(ctx) {
		d := event.NewDefaultOperateEvent()
		if err := resp.Decode(d); err != nil {
			return nil, exception.NewInternalServerError("decode event error, error is %s", err)
		}

		set.Add(d)
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get event count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}
