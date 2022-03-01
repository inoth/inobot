local gohttp = require("gohttp")
local util = require("script.package.util")

http = {}

function http.get(url,params,header)
    local param = util.SerializedGetParams(params)
    print("get 请求参数: " .. param)
    resp,ok = gohttp.get(url,param,header)
    if ok then
        return resp
    end
    return nil
end

function http.post(url,params,header)
    local param = util.SerializedPostParams(params)
    resp,ok = gohttp.post(url,param,header)
    if ok then
        return resp
    end
    return nil
end

return http