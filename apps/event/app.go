package event

import (
	"encoding/json"
	"time"

	request "github.com/infraboard/mcube/v2/http/request"
	"github.com/rs/xid"
	"google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	AppName = "event"
)

// 事件主题定义(由事件类型确定)
// 1. 操作事件

func NewJsonOperateEvent(e *OperateEventData) (*Event, error) {
	return NewOperateEvent(ContentType_JSON, e)
}

func NewProtoOperateEvent(e *OperateEventData) (*Event, error) {
	return NewOperateEvent(ContentType_PROTOBUF, e)
}

// NewOperateEvent 实例
func NewOperateEvent(t ContentType, e *OperateEventData) (*Event, error) {
	var err error

	obj := &Event{
		Id:     xid.New().String(),
		Type:   Type_OPERATE,
		Header: NewHeader(),
		Body:   new(anypb.Any),
	}
	obj.Header.ContentType = t

	switch t {
	case ContentType_JSON:
		obj.Body.Value, err = json.Marshal(e)
	default:
		obj.Body, err = anypb.New(e)
	}

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func NewDefaultOperateEvent() *OperateEvent {
	return &OperateEvent{
		Header: NewHeader(),
	}
}

// NewDefaultEvent todo
func NewDefaultEvent() *Event {
	return &Event{
		Header: NewHeader(),
	}
}

// NewHeader todo
func NewHeader() *Header {
	return &Header{
		Time:        time.Now().UnixMilli(),
		Meta:        make(map[string]string),
		ContentType: ContentType_PROTOBUF,
	}
}

// DecodeBody 解码body数据
func (e *Event) ParseBoby(body proto.Message) (err error) {
	switch e.Header.ContentType {
	case ContentType_JSON:
		err = json.Unmarshal(e.Body.Value, body)
	default:
		err = anypb.UnmarshalTo(e.Body, body, proto.UnmarshalOptions{})
	}
	return err
}

// Validate 校验事件是否合法
func (e *Event) Validate() error {
	return nil
}

// GetMetaKey 获取meta信息
func (e *Event) GetMetaKey(key string) (string, bool) {
	v, ok := e.Header.Meta[key]
	return v, ok
}

// SetMeta 设置meta信息
func (e *Event) SetMeta(key, value string) {
	e.Header.Meta[key] = value
}

// SetLevel 设置事件级别
func (e *Event) SetLevel(l Level) {
	e.Header.Level = l
}

// SetSource 设置事件来源
func (e *Event) SetSource(src string) {
	e.Header.Source = src
}

// NewQueryEventkRequest 查询book列表
func NewQueryEventkRequest(page *request.PageRequest) *QueryEventRequest {
	return &QueryEventRequest{
		Page: page,
	}
}

// NewSaveEventRequest todo
func NewSaveEventRequest() *SaveEventRequest {
	return &SaveEventRequest{}
}

func (req *SaveEventRequest) Add(item *Event) {
	req.Items = append(req.Items, item)
}

func (req *SaveEventRequest) Ids() []string {
	ids := make([]string, 0, len(req.Items))
	for i := range req.Items {
		ids = append(ids, req.Items[i].Id)
	}

	return ids
}

func (req *SaveEventRequest) ParseEvent() ([]interface{}, error) {
	docs := make([]interface{}, 0, len(req.Items))
	for i := range req.Items {
		switch req.Items[i].Type {
		case Type_OPERATE:
			data := &OperateEventData{}
			err := req.Items[i].ParseBoby(data)
			if err != nil {
				return nil, err
			}
			oe := &OperateEvent{
				Id:     req.Items[i].Id,
				SaveAt: time.Now().UnixMilli(),
				Type:   req.Items[i].Type,
				Header: req.Items[i].Header,
				Body:   data,
			}
			if err != nil {
				return nil, err
			}
			docs = append(docs, oe)
		}
	}

	return docs, nil
}

// NewSaveReponse todo
func NewSaveReponse() *SaveReponse {
	return &SaveReponse{}
}

func (resp *SaveReponse) AddSuccess(ids ...string) {
	for i := range ids {
		resp.Success = append(resp.Success, ids[i])
	}
}

func (resp *SaveReponse) AddFailed(ids ...string) {
	for i := range ids {
		resp.Failed = append(resp.Failed, ids[i])
	}
}

// NewOperateEventSet 实例
func NewOperateEventSet() *OperateEventSet {
	return &OperateEventSet{
		Items: []*OperateEvent{},
	}
}

func (s *OperateEventSet) Add(item *OperateEvent) {
	s.Items = append(s.Items, item)
}
