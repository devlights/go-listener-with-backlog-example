# go-listener-with-backlog-example

TCP Listener with Backlog examples by golang

![Go Version](https://img.shields.io/badge/go-1.19-blue.svg)

## Go version

```shell script
$ lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 20.04.5 LTS
Release:        20.04
Codename:       focal

$ go version
go version go1.19.4 linux/amd64
```

## How to Run

### Using [go-task](https://taskfile.dev/)

```sh
$ task run
task: [run] go run .
[With Context] 07:56:16 start
[With Context] 07:56:17 dial tcp4 127.0.0.1:33333: i/o timeout
[With Context] 07:56:17 done
```

## REFERENCES

- https://github.com/golang/go/issues/39000
- https://github.com/golang/go/issues/41470
- https://github.com/valyala/tcplisten/blob/master/tcplisten.go
- https://stackoverflow.com/a/49593356
- https://stackoverflow.com/a/46279623
- https://meetup-jp.toast.com/1509
