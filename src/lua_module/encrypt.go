package luamodule

import lua "github.com/yuin/gopher-lua"

func LoadEncryptModule(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), encryptExports)
	// returns the module
	L.Push(mod)
	return 1
}

var encryptExports = map[string]lua.LGFunction{
	"md5": producer,
}

func md5(l *lua.LState) int {
	return 1
}
