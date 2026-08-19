// Copyright IBM Corp. 2016, 2026
// SPDX-License-Identifier: MIT

// +build linux darwin dragonfly freebsd netbsd openbsd solaris

package gsyslog

import (
	"net"
	"testing"
	"time"
)

func TestDialLoggerWithOptionsLocal(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			c.Close()
		}
	}()

	l, err := DialLoggerWithOptions("tcp", ln.Addr().String(), LOG_INFO, "USER", "gsyslog-test", DialOptions{Timeout: time.Second})
	if err != nil {
		t.Fatal(err)
	}
	if err := l.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestDialLoggerZeroTimeoutUsesDefault(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			c.Close()
		}
	}()

	l, err := DialLogger("tcp", ln.Addr().String(), LOG_INFO, "USER", "gsyslog-test")
	if err != nil {
		t.Fatal(err)
	}
	if err := l.Close(); err != nil {
		t.Fatal(err)
	}
}
