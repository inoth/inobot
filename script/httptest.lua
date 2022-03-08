local http = require("script.lib.http")

-- go调用时注入的公共参数
-- local callbacktype 回调类型，http/nsq（未实现）
-- local callback 回调地址，url/topic
-- local args 请求包含的参数内容

-- 加上这玩防止命名不小心重复
local script_handler = {}

function script_handler.main()
    -- 打印注入参数内容
    print(callback)

    local http_header = {
        Authorization = "xxxxxx"
    }
    local resp,ok = http.get(args.url,args.body,http_header)
    if not ok then
        print("接口请求失败")
    end
    print("脚本请求接口返回： ")
    for i, v in pairs(resp) do
        print(i..": ",v)
    end

    -- 回调
    script_handler.callback(resp)
end

function script_handler.callback(respData)
    local resp,ok = http.post(callback,respData)
    if not ok then
        print("回调请求失败")
    end
    print("脚本回调请求返回： ")
    for i, v in pairs(resp) do
        print(i..": ",v)
    end
end

script_handler.main()