# sandbox
This tool generates sandbox [Go](https://go.dev) project in temporary files directory and opens it in editor. An editor from [`$EDITOR`](https://askubuntu.com/a/432530) environment variable is used, if set. Otherwise [Visual Studio Code](https://code.visualstudio.com)'s `code` command is used by default.

Generated project consists of two files:
```
├── go.mod
└── main.go
```
[`go.mod`](https://go.dev/doc/tutorial/create-module) file is created by tool using `go mod init sandbox` command (so module is of a version of Go toolchain installed in your system).
## Install
With a [correctly configured](https://go.dev/doc/install#testing) Go toolchain run
```
go install github.com/electrofocus/sandbox@latest
```
## Use
To open editor run
```
sandbox
```