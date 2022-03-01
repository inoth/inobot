package executor

import (
	"encoding/json"
	"fmt"

	luamodule "github.com/inoth/inobot/src/lua_module"
	lua "github.com/yuin/gopher-lua"
)

type LuaExecutor struct{}

func (LuaExecutor) Processing(body []byte) error {
	// var data consumer.LuaScriptMessageBody
	// if err := json.Unmarshal(body, &data); err != nil {
	// 	fmt.Println("参数解析错误")
	// 	return nil
	// }
	fmt.Println("开始执行脚本")
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("参数解析错误")
		return nil
	}
	// TODO:使用连接池优化脚本加载
	l := luaPool.Get().(*lua.LState)
	defer luaPool.Put(l)
	// 调起加载lua脚本执行
	l.SetGlobal("args", luamodule.MapToTable(data))
	// l.DoFile(data.ScriptPath)
	err := l.DoFile("script/test.lua")
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	return nil
}
