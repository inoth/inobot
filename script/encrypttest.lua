local encrypt = require("script.lib.encrypt")

-- go调用时注入的公共参数
-- local callbacktype 回调类型，http/nsq（未实现）
-- local callback 回调地址，url/topic
-- local args 请求包含的参数内容

-- 加上这玩防止命名不小心重复
local encrypt_handler = {}

function encrypt_handler.main()
    local data = "abcdefg"
    print("md5：" .. encrypt.md5(data))
    
    local sign = encrypt.aes_encrypt(data)
    print("aes encrypt:" .. sign)
    print("aes decrypt:" .. encrypt.aes_decrypt(sign))
end


encrypt_handler.main()