local http = require("script.lib.http")

-- go调用时注入的公共参数
-- local callbacktype 回调类型，http/nsq（未实现）
-- local callback 回调地址，url/topic
-- local args 请求包含的参数内容

function main()
    local http_header = {
        Authorization = "someusertoken"    
    }
    resp,ok = http.get(args.url,args.body,http_header)
    if not ok then
        print("接口请求失败")
    end
    print("脚本请求接口返回： ")
    for i, v in pairs(resp) do
        print(i..": ",v)
    end

    -- 回调
    callback(resp)
end

function callback(respData)
    resp,ok = http.post(callback,respData)
    if not ok then
        print("回调请求失败")
    end
    print("脚本回调请求返回： ")
    for i, v in pairs(resp) do
        print(i..": ",v)
    end
end

main()