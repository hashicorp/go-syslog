// Copyright IBM Corp. 2016, 2026
// SPDX-License-Identifier: MIT

package gsyslog

import "time"

// DialOptions holds optional settings for DialLoggerWithOptions.
type DialOptions struct {
	// Timeout is the remote dial deadline. A zero value keeps the
	// existing 50 millisecond default.
	Timeout time.Duration
}

// Priority maps to the syslog priority levels
type Priority int

const (
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

// Syslogger interface is used to write log messages to syslog
type Syslogger interface {
	// WriteLevel is used to write a message at a given level
	WriteLevel(Priority, []byte) error

	// Write is used to write a message at the default level
	Write([]byte) (int, error)

	// Close is used to close the connection to the logger
	Close() error
}
