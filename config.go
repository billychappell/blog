package main

import (
	"net"
	"strings"
	"time"
)

func httpsHost() string {
	if *prod {
		return "" // TODO: Fill in prod address
	}
	if v := *httpAddr; strings.HasPrefix(v, ":") {
		return "localhost" + v
	} else {
		return v
	}
}

func httpHost() string {
	if *prod {
		return "" // TODO: Fill in prod address
	}
	if v := *httpAddr; strings.HasPrefix(v, ":") {
		return "localhost" + v
	} else {
		return v
	}
}

type tcpKeepAlive struct {
	*net.TCPListener
}

func (ln tcpKeepAlive) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
