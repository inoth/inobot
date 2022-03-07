package luamodule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

func LoadHttpModule(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), httpExports)
	// returns the module
	L.Push(mod)
	return 1
}

var httpExports = map[string]lua.LGFunction{
	"get":  get,
	"post": post,
}

func get(l *lua.LState) int {
	uri := l.ToString(1)
	params := l.ToString(2)
	header := l.ToTable(3)
	url := fmt.Sprintf("%s?%s", uri, params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LFalse)
		return 2
	}
	header.ForEach(func(l1, l2 lua.LValue) {
		fmt.Printf("key = %s;val = %s\n", l1.String(), l2.String())
		req.Header.Set(l1.String(), l2.String())
	})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("%s 请求失败：%s", url, err.Error())
		l.Push(lua.LNil)
		l.Push(lua.LFalse)
		return 2
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	buf, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(buf))
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LFalse)
		return 2
	}
	err = json.Unmarshal(buf, &data)
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LFalse)
		return 2
	}
	l.Push(MapToTable(data))
	l.Push(lua.LTrue)
	return 2
}

func post(l *lua.LState) int {
	url := l.ToString(1)
	params := l.ToString(2)
	header := l.ToTable(3)
	logrus.Infof("脚本post请求参数： %v", params)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(params))
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LFalse)
		return 2
	}
	header.ForEach(func(l1, l2 lua.LValue) {
		fmt.Printf("key = %s;val = %s\n", l1.String(), l2.String())
		req.Header.Set(l1.String(), l2.String())
	})
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("%s 请求失败：%s", url, err.Error())
		logrus.Errorf("%s 请求失败参数：%s", params, err.Error())
		l.Push(lua.LNil)
		l.Push(lua.LFalse)
		return 2
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LFalse)
		return 2
	}
	err = json.Unmarshal(buf, &data)
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LFalse)
		return 2
	}
	l.Push(MapToTable(data))
	l.Push(lua.LTrue)
	return 2
}
