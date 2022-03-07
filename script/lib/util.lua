local json = require("script.lib.json")

util = {} 

function util.SerializedGetParams(params)
    local param = ""
    for k, v in pairs(params) do
        param = param .. k .. "=" .. v .. "&"
    end
    return string.sub(param,1,-2)
end

function util.SerializedPostParams(params)
    local str = json.encode(params)
    return str
end

return util