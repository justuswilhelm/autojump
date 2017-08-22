function cd
    if builtin cd $argv
        autojump --record --location (realpath "$PWD")
    end
end
