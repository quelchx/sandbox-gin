# Instructions on how to setup Golang

## Setup

Create a directory like so from the `go` directory: `/src/github.com/quelchx/project-name` then run `go mod init` to initialize the project.

Will need to have `CompileDaemon` installed to run the project. To install run `go get github.com/githubnemo/CompileDaemon` and `go install -mod=mod github.com/githubnemo/CompileDaemon`

After installing golang ensure the .zshrc file is as follows (macOS):

```sh
# If you come from bash you might have to change your $PATH.
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
