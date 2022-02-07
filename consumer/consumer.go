package consumer

import (
	"github.com/nsqio/go-nsq"
)

type ConsumerConfig struct {
	Addr     string
	Channel  string
	Topic    string
	PoolSize int
}

type IConsumer interface {
	// Init(addr, port, channel, topic string, poolSize int, scheduler *script.ScriptScheduler) error
	GetConfig() ConsumerConfig
}

func Run(consumer IConsumer) error {
	conf := consumer.GetConfig()
	cfg := nsq.NewConfig()
	for i := 0; i < conf.PoolSize; i++ {
		go func() {
			c, err := nsq.NewConsumer(conf.Topic, conf.Channel, cfg)
			if err != nil {
				panic(err)
			}

			c.AddHandler(consumer.(nsq.Handler))

			if err := c.ConnectToNSQD(conf.Addr); err != nil {
				panic(err)
			}
		}()
	}
	return nil
}

type MessageBody struct {
	// 执行命令
	Cmd string `json:"cmd"`
	// 脚本源
	ScriptSource string `json:"scriptSource"`
	// 脚本地址
	Script string `json:"script"`
	// 导入脚本参数
	Args string `json:"args"`
}
