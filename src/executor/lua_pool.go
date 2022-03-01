package executor

import (
	"fmt"
	"sync"

	luamodule "github.com/inoth/inobot/src/lua_module"
	lua "github.com/yuin/gopher-lua"
)

var luaPool *sync.Pool

type LuaPool struct{}

func (LuaPool) Init() error {
	luaPool = &sync.Pool{
		New: func() interface{} {
			L := lua.NewState()
			// init global function
			// 装载加密模块
			// 装载http请求模块
			fmt.Println("装载luahttp模块。。。")
			L.PreloadModule("gohttp", luamodule.LoadHttpModule)
			return L
		},
	}
	return nil
}
