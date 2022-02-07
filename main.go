package main

import (
	"github.com/inoth/inobot/consumer"
	"github.com/inoth/inobot/proxy"
	"github.com/inoth/inobot/script"
)

func init() {
	proxy.AutoSet()
}

func main() {
	script.Init(10)
	consumer.NsqInit("localhost:4150", "ngx_body_mid", "ngx_body_mid", 10)

	go script.CallScheduler().Run()

	if err := consumer.Run(consumer.GetNsqConsumer()); err != nil {
		panic(err)
	}

	<-make(chan struct{})
}
