local http = require("script.package.http")

t = {
    Authorization = "xxxxxxxx"    
}

function main()
    resp,ok = http.get("http://localhost:8080",{aaa="111",bbb = "222"},t)
    if ok then
        for i, v in pairs(resp) do
            print(i..": ",v)
        end
    end
end 

main()