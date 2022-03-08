local goencrypt = require("goencrypt")

encrypt = {}

function encrypt.md5(str)
    return goencrypt.md5(str)
end

function encrypt.aes_encrypt(str,key)
    key = key or ""
    local data,ok = goencrypt.aesEncrypt(str,key)
    if not ok then
        print("加密失败:" .. str)
        return data
    end
    return data
end

function encrypt.aes_decrypt(str,key)
    key = key or ""
    local data,ok = goencrypt.aesDecrypt(str,key)
    if not ok then
        print("解密失败:" .. str)
        return data
    end
    return data
end

return encrypt