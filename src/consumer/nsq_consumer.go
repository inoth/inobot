package consumer

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

type LuaScriptWorkers interface {
	// Before() error
	Processing(body []byte) error
	// After() error
}

type LuaScriptMessageBody struct {
	ScriptPath string                 // 脚本地址
	Args       map[string]interface{} // 脚本参数
}

type LuaScriptExec struct {
	Topic   string           // 订阅key
	Channel string           // 订阅渠道
	Workers LuaScriptWorkers // 消息处理执行
}

func (c *LuaScriptExec) GetSubscription() (string, string) {
	return c.Topic, c.Channel
}

func (c *LuaScriptExec) HandleMessage(msg *nsq.Message) error {
	fmt.Println(c.Topic)
	fmt.Println(c.Channel)
	fmt.Println(string(msg.Body))
	return c.Workers.Processing(msg.Body)
}
