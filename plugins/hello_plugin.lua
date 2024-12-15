function HandleShortcode(arguments)

    -- title is 1
    -- content is 2
    --  excerpt is 3
    local title = arguments[1]
    local content = arguments[2]
    local excerpt = arguments[3]

    post = post.new()
    post:title()
    post:content()
    post:excerpt()

    return post
end