package consumer

import (
	"fmt"

	"github.com/inoth/inobot/script"
	"github.com/nsqio/go-nsq"
)

var nc *NsqConsumer

type NsqConsumer struct {
	Config ConsumerConfig
}

func NsqInit(addr, channel, topic string, poolSize int) error {
	nc = &NsqConsumer{
		Config: ConsumerConfig{
			Addr:     addr,
			Channel:  channel,
			Topic:    topic,
			PoolSize: poolSize,
		},
	}
	return nil
}

func GetNsqConsumer() IConsumer {
	return nc
}

func (c *NsqConsumer) GetConfig() ConsumerConfig {
	return c.Config
}

func (c *NsqConsumer) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("消费者预处理：%v\n", string(msg.Body))
	script.CallScheduler().SetQeust(string(msg.Body))
	return nil
}
