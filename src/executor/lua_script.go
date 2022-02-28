package executor

import (
	"encoding/json"
	"fmt"

	"github.com/inoth/inobot/src/consumer"
)

type LuaExecutor struct {
	// DOTO:使用对象池优化加载脚本速度
}

func Processing(body []byte) error {
	var data consumer.LuaScriptMessageBody
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("参数解析错误")
		return nil
	}
	// TODO:使用连接池优化脚本加载
	// 调起加载lua脚本执行
	return nil
}
