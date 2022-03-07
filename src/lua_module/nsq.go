package luamodule

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func LoadNsqModule(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), httpExports)
	// returns the module
	L.Push(mod)
	return 1
}

var nsqExports = map[string]lua.LGFunction{
	"producer": producer,
}

func producer(l *lua.LState) int {
	data := l.ToString(1)
	fmt.Println(data)
	return 1
}
