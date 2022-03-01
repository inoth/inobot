local json = require("script.package.dkjson")

util = {} 

function util.SerializedGetParams(params)
    local param = ""
    for k, v in pairs(params) do
        param = param .. k .. "=" .. v .. "&"
    end
    return string.sub(param,1,-2)
end

function util.SerializedPostParams(params)
    local str = json.encode(params, { indent = true })
    return st
end

return util