package impl_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/infraboard/maudit/apps/event"
	"github.com/segmentio/kafka-go"
)

func TestPublishEvent(t *testing.T) {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "audit_event",
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	data := &event.OperateEventData{
		UserName:     "张三",
		ServiceName:  "cmdb",
		ResourceType: "secret",
		Action:       "create",
		Request:      "{}",
		Response:     "{}",
	}
	e, err := event.NewJsonOperateEvent(data)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Value: payload,
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func TestConsumerEvent(t *testing.T) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group-id",
		Topic:    "audit_event",
		MaxBytes: 10e6, // 10MB,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		// 处理完消息后需要提交该消息已经消费完成, 消费者挂掉后保存消息消费的状态
		// FetchMessage() / CommitMessages(ctx, m) 分段提交
		// if err := r.CommitMessages(ctx, m); err != nil {
		//     log.Fatal("failed to commit messages:", err)
		// }
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
