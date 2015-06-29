# colorgo
`colorgo` is a wrapper to `go` command that colorizes output from `go build` and `go test`.
![screenshot 1](http://songgao.github.com/colorgo/images/screenshot1.png)

# Installation
```
go get -u github.com/songgao/colorgo
```

# Usage
```bash
colorgo build
```

# alias
`colorgo` changes nothing to sub-commands other than `go build`. So you can optionally define alias in your shell conf so that `go build` always prints colorized error message:

bash: `~/.bashrc`
```
alias go=colorgo
```

fish-shell: `~/.config/fish/config.fish`
```
alias go colorgo
```

# License
[BSD 3-Clause License](http://opensource.org/licenses/BSD-3-Clause)
