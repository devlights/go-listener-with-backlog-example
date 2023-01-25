package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var (
		proto   = "tcp4"
		addr    = net.JoinHostPort("localhost", "33333")
		backlog = 1
	)

	listener, err := Listen(addr, backlog)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// Fill up the backlog to time out
	net.Dial(proto, addr)
	net.Dial(proto, addr)

	// Connect
	var (
		connLog  = log.New(os.Stdout, "[With Context] ", log.Ltime)
		ctx, cxl = context.WithTimeout(context.Background(), 1*time.Second)
		dialer   = net.Dialer{}
	)
	defer cxl()

	connLog.Println("start")
	{
		_, err = dialer.DialContext(ctx, proto, addr)
		if err != nil {
			connLog.Println(err)
		}
	}
	connLog.Println("done")
}
