// Copyright IBM Corp. 2016, 2026
// SPDX-License-Identifier: MIT

// +build windows plan9 nacl

package gsyslog

import (
	"fmt"
)

// NewLogger is used to construct a new Syslogger
func NewLogger(p Priority, facility, tag string) (Syslogger, error) {
	return nil, fmt.Errorf("Platform does not support syslog")
}

// DialLogger is used to construct a new Syslogger that establishes connection to remote syslog server
func DialLogger(network, raddr string, p Priority, facility, tag string) (Syslogger, error) {
	return DialLoggerWithOptions(network, raddr, p, facility, tag, DialOptions{})
}

func DialLoggerWithOptions(network, raddr string, p Priority, facility, tag string, opts DialOptions) (Syslogger, error) {
	return nil, fmt.Errorf("Platform does not support syslog")
}
