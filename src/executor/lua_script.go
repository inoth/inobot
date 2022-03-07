package executor

import (
	"github.com/inoth/inobot/src/consumer"
	luamodule "github.com/inoth/inobot/src/lua_module"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

type LuaExecutor struct{}

func (LuaExecutor) Processing(body []byte) error {
	// var data consumer.LuaScriptMessageBody
	// if err := json.Unmarshal(body, &data); err != nil {
	// 	fmt.Println("参数解析错误")
	// 	return nil
	// }
	// fmt.Println("开始执行脚本")
	// var data map[string]interface{}
	// if err := json.Unmarshal(body, &data); err != nil {
	// 	fmt.Println("参数解析错误")
	// 	return nil
	// }

	data := consumer.LuaScriptMessageBody{
		CallBackType: "http",
		CallBack:     "http://localhost:8080/callback",
		Args: map[string]interface{}{
			"method": "get",
			"url":    "http://localhost:8080",
			"body": map[string]interface{}{
				"id":   1,
				"name": "inoth",
			},
		},
		ScriptPath: "script/httptest.lua",
	}
	// TODO:使用连接池优化脚本加载
	l := luaPool.Get().(*lua.LState)
	defer luaPool.Put(l)
	// 调起加载lua脚本执行
	l.SetGlobal("callbacktype", lua.LString(data.CallBackType))
	l.SetGlobal("callback", lua.LString(data.CallBack))
	l.SetGlobal("args", luamodule.MapToTable(data.Args))
	err := l.DoFile(data.ScriptPath)
	if err != nil {
		logrus.Errorf("%s 调用失败：%s", data.ScriptPath, err.Error())
		return err
	}
	return nil
}
