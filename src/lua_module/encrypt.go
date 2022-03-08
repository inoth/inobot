package luamodule

import (
	"github.com/inoth/inobot/utils"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

func LoadEncryptModule(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), encryptExports)
	// returns the module
	L.Push(mod)
	return 1
}

var encryptExports = map[string]lua.LGFunction{
	"md5":        md5,
	"aesEncrypt": aesEncrypt,
	"aesDecrypt": aesDecrypt,
}

func md5(l *lua.LState) int {
	str := l.ToString(1)
	sign := utils.Md5(str)
	l.Push(lua.LString(sign))
	return 1
}

func aesEncrypt(l *lua.LState) int {
	str := l.ToString(1)
	key := l.ToString(2)
	if len(key) > 0 {
		sign, err := utils.Encrypt(key, str)
		if err != nil {
			logrus.Errorf("加密失败：%s；key：%s;原串：%s", err.Error(), key, str)
			l.Push(lua.LString(""))
			l.Push(lua.LFalse)
			return 2
		}
		l.Push(lua.LString(sign))
		l.Push(lua.LTrue)
	} else {
		sign, ok := utils.DefukaEncrypt(str)
		if !ok {
			logrus.Errorf("默认key加密失败，原串：%s", str)
			l.Push(lua.LString(""))
			l.Push(lua.LFalse)
			return 2
		}
		l.Push(lua.LString(sign))
		l.Push(lua.LTrue)
	}
	return 2
}

func aesDecrypt(l *lua.LState) int {
	str := l.ToString(1)
	key := l.ToString(2)
	if len(key) > 0 {
		origin, err := utils.Decrypt(key, str)
		if err != nil {
			logrus.Errorf("解密失败：%s；key：%s;原串：%s", err.Error(), key, str)
			l.Push(lua.LString(""))
			l.Push(lua.LFalse)
			return 2
		}
		l.Push(lua.LString(origin))
		l.Push(lua.LTrue)
	} else {
		origin, ok := utils.DefukaDecrypt(str)
		if !ok {
			logrus.Errorf("默认key解密失败，原串：%s", str)
			l.Push(lua.LString(""))
			l.Push(lua.LFalse)
			return 2
		}
		l.Push(lua.LString(origin))
		l.Push(lua.LTrue)
	}
	return 2
}
