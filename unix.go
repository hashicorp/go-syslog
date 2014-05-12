// +build linux darwin freebsd openbsd solaris

package gsyslog

import (
	"fmt"
	"log/syslog"
)

// builtinLogger wraps the Golang implementation of a
// syslog.Writer to provide the Syslogger interface
type builtinLogger struct {
	*syslog.Writer
}

// NewLogger is used to construct a new Syslogger
func NewLogger(p Priority, tag string) (Syslogger, error) {
	priority := syslog.Priority(p) | syslog.LOG_DAEMON
	l, err := syslog.New(priority, tag)
	if err != nil {
		return nil, err
	}
	return &builtinLogger{l}, nil
}

// WriteLevel writes out a message at the given priority
func (b *builtinLogger) WriteLevel(p Priority, buf []byte) error {
	switch p {
	case LOG_EMERG:
		return b.Emerg(string(buf))
	case LOG_ALERT:
		return b.Alert(string(buf))
	case LOG_CRIT:
		return b.Crit(string(buf))
	case LOG_ERR:
		return b.Err(string(buf))
	case LOG_WARNING:
		return b.Warning(string(buf))
	case LOG_NOTICE:
		return b.Notice(string(buf))
	case LOG_INFO:
		return b.Info(string(buf))
	case LOG_DEBUG:
		return b.Debug(string(buf))
	default:
		return fmt.Errorf("Unknown priority: %v", p)
	}
}
