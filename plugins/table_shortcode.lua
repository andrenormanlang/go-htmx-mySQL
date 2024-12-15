-- User has to know what "arguments" will look like
-- from a callback perspective. I.e. how many strings it has,
-- the type of each string in a position, etc.
function HandleShortcode(arguments)
    if #arguments ~= 2 then
        return ""
    end

    return "\n| Tables   |      Are      |  Cool |\n|----------|:-------------:|------:|\n| col 1 is |  left-aligned | $1600 |\n| col 2 is |    centered   |   $12 |\n| col 3 is | right-aligned |    $1 |\n"
end