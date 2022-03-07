-- local json = require("script.lib.dkjson")
local cjson = require "cjson"

util = {} 

function util.SerializedGetParams(params)
    local param = ""
    for k, v in pairs(params) do
        param = param .. k .. "=" .. v .. "&"
    end
    return string.sub(param,1,-2)
end

function util.SerializedPostParams(params)
    local str = cjson.encode(params)
    return str
end

return util