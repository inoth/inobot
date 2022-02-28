function main()
    print(args.method)
    for k, v in pairs(args.body) do
        print(k..": ",v)
    end
end 

main()