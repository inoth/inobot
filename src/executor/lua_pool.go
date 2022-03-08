package executor

import (
	"sync"

	luamodule "github.com/inoth/inobot/src/lua_module"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

var luaPool *sync.Pool

type LuaPool struct{}

func (LuaPool) Init() error {
	luaPool = &sync.Pool{
		New: func() interface{} {
			L := lua.NewState()
			logrus.Info("装载luahttp模块...")
			L.PreloadModule("gohttp", luamodule.LoadHttpModule)
			logrus.Info("装载luaencrypt模块...")
			L.PreloadModule("goencrypt", luamodule.LoadEncryptModule)
			return L
		},
	}
	return nil
}
