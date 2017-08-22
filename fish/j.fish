function j
    set -l fuzz "$argv[1]"
    if test -z "$fuzz"
        echo "Usage: j DESTINATION"
        return 1
    end
    set -l target (autojump --location "$fuzz")
    if test -z "$target"
        echo "Could not find $fuzz"
        return 1
    end
    cd "$target"
end
