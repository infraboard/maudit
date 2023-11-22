package impl

import (
	"context"
	"fmt"
)

func (i *impl) ConsumerEvent() {
	for {
		m, err := i.kr.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
