package queue

import (
	"fmt"
	"os"

	"github.com/inoth/inobot/components/config"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

type QueueHandler interface {
	GetSubscription() (string, string)
}

type NsqQueue struct {
	Host string
	// Consumers []*nsq.Handler
	consumers []QueueHandler
}

func (nq *NsqQueue) AddConsumer(consumers ...QueueHandler) *NsqQueue {
	nq.consumers = consumers[:]
	return nq
}

func (nq *NsqQueue) ServeStart() error {
	replica := config.Cfg.GetInt("replica")
	for _, consumer := range nq.consumers {
		for i := 0; i < int(replica); i++ {
			logrus.Info("消费者初始化...")
			go func() {
				defer func() {
					if exception := recover(); exception != nil {
						if err, ok := exception.(error); ok {
							fmt.Printf("%v\n", err)
						} else {
							panic(exception)
						}
						os.Exit(1)
					}
				}()
				topic, channel := consumer.GetSubscription()
				c, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
				if err != nil {
					panic(err)
				}
				c.AddHandler(consumer.(nsq.Handler))
				if err := c.ConnectToNSQD(nq.Host); err != nil {
					panic(err)
				}
			}()
		}
	}
	<-make(chan struct{})
	// select{}
	return nil
}
