# Autojump

Record locations you `cd` to, then lets you jump to the same location using
a fuzzy search string.

## Example

```
cd /etc/
pwd  # will output /etc/
cd $HOME
autojump tc
pwd  # will output /etc/
```

## Install

```
go get -u github.com/justuswilhelm/autojump
```

### Fish

```
ln -fv fish/j.fish $XDG_CONFIG_HOME/fish/functions
ln -fv fish/cd.fish $XDG_CONFIG_HOME/fish/functions
```
