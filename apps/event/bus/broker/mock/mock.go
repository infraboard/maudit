package broker

import (
	"fmt"

	"github.com/infraboard/maudit/apps/event"
	"github.com/infraboard/maudit/apps/event/bus"
)

// NewBroker todo
func NewBroker() bus.Publisher {
	return &mockBroker{}
}

type mockBroker struct {
}

func (m *mockBroker) Pub(topic string, e *event.Event) error {
	fmt.Println(topic, e)
	return nil
}
