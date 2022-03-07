local gohttp = require("gohttp")
local util = require("script.lib.util")

http = {}

function http.get(url,params,header)
    header = header or {}
    local param = util.SerializedGetParams(params)
    print("get 请求参数: " .. param)
    return gohttp.get(url,param,header)
end

function http.post(url,params,header)
    header = header or {} 
    local param = util.SerializedPostParams(params)
    print("post 请求地址: " .. url)
    print("post 请求参数: " .. param)
    return gohttp.post(url,param,header)
end

return http