local gohttp = require("gohttp")
local util = require("script.lib.util")

http = {}

function http.get(url,params,header)
    header = header or {}
    local param = util.SerializedGetParams(params)
    print("get 请求参数: " .. param)
    return gohttp.get(url,param,header)
    -- if not ok then
    --     print("接口请求失败")
    --     return nil
    -- end
    -- return resp
end

function http.post(url,params,header)
    header = header or {}
    local param = util.SerializedPostParams(params)
    return gohttp.post(url,param,header)
    -- if not ok then
    --     print("接口请求失败")
    --     return nil
    -- end
    -- return resp
end

return http