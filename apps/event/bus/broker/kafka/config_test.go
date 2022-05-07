package kafka_test

import (
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/maudit/apps/event/bus/broker/kafka"
)

var (
	testConfig = `hosts:
  - "192.168.100.1:9092"`
)

func TestConfig(t *testing.T) {
	should := assert.New(t)
	conf := kafka.NewDefultConfig()
	if should.NoError(yaml.Unmarshal([]byte(testConfig), conf)) {
		should.Equal(conf.Hosts[0], "192.168.100.1:9092")
	}
}
