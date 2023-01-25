package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

// listen は、指定された情報で net.Listener を生成して返します.
//
// Go の net.Listen は、バックログが指定できないため
// テスト目的でバックログを極端に少ない状態にすることが出来ない。
// 利用される値は linux の場合は net.core.somaxconn の値であり
// 存在しない場合は、 syscall.SOMAXCONN の定数値が利用される。
//
// なので、バックログを指定したい場合はソケットを直接生成して
// net.FileListener に紐づけて、リスナーとして利用する。
//
// # REFERENCES
//   - https://github.com/golang/go/issues/39000
//   - https://github.com/golang/go/issues/41470
//   - https://github.com/valyala/tcplisten/blob/master/tcplisten.go
//   - https://stackoverflow.com/a/49593356
//   - https://stackoverflow.com/a/46279623
//   - https://meetup-jp.toast.com/1509
func Listen(addr string, backLog int) (net.Listener, error) {
	// make tcp addr
	var (
		tcpAddr *net.TCPAddr
		err     error
	)

	tcpAddr, err = net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		return nil, err
	}

	// make socket addr
	var (
		sockAddr syscall.SockaddrInet4
	)

	sockAddr.Port = tcpAddr.Port
	copy(sockAddr.Addr[:], tcpAddr.IP.To4())

	// make socket file descriptor
	var (
		sockFd     int
		sockDomain = syscall.AF_INET
		sockType   = syscall.SOCK_STREAM | syscall.SOCK_NONBLOCK | syscall.SOCK_CLOEXEC
		sockProto  = syscall.IPPROTO_TCP
	)

	sockFd, err = syscall.Socket(sockDomain, sockType, sockProto)
	if err != nil {
		return nil, err
	}

	// bind
	err = syscall.Bind(sockFd, &sockAddr)
	if err != nil {
		return nil, err
	}

	// listen
	err = syscall.Listen(sockFd, backLog)
	if err != nil {
		return nil, err
	}

	// make net.Listener
	var (
		fname    = fmt.Sprintf("backlog.%d.%s.%s", os.Getpid(), "tcp4", addr)
		file     = os.NewFile(uintptr(sockFd), fname)
		listener net.Listener
	)
	defer file.Close()

	listener, err = net.FileListener(file)
	if err != nil {
		return nil, err
	}

	return listener, nil
}
