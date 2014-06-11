// +build linux darwin freebsd openbsd solaris

package gsyslog

import (
	"fmt"
	"log/syslog"
	"strings"
)

// builtinLogger wraps the Golang implementation of a
// syslog.Writer to provide the Syslogger interface
type builtinLogger struct {
	*syslog.Writer
}

// NewLogger is used to construct a new Syslogger
func NewLogger(p Priority, facility, tag string) (Syslogger, error) {
	fPriority, err := facilityPriority(facility)
	if err != nil {
		return nil, err
	}
	priority := syslog.Priority(p) | fPriority
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

// facilityPriority converts a facility string into
// an appropriate priority level or returns an error
func facilityPriority(facility string) (syslog.Priority, error) {
	facility = strings.ToUpper(facility)
	switch facility {
	case "KERN":
		return syslog.LOG_KERN, nil
	case "USER":
		return syslog.LOG_USER, nil
	case "MAIL":
		return syslog.LOG_MAIL, nil
	case "DAEMON":
		return syslog.LOG_DAEMON, nil
	case "AUTH":
		return syslog.LOG_AUTH, nil
	case "SYSLOG":
		return syslog.LOG_SYSLOG, nil
	case "LPR":
		return syslog.LOG_LPR, nil
	case "NEWS":
		return syslog.LOG_NEWS, nil
	case "UUCP":
		return syslog.LOG_UUCP, nil
	case "CRON":
		return syslog.LOG_CRON, nil
	case "AUTHPRIV":
		return syslog.LOG_AUTHPRIV, nil
	case "FTP":
		return syslog.LOG_FTP, nil
	case "LOCAL0":
		return syslog.LOG_LOCAL0, nil
	case "LOCAL1":
		return syslog.LOG_LOCAL1, nil
	case "LOCAL2":
		return syslog.LOG_LOCAL2, nil
	case "LOCAL3":
		return syslog.LOG_LOCAL3, nil
	case "LOCAL4":
		return syslog.LOG_LOCAL4, nil
	case "LOCAL5":
		return syslog.LOG_LOCAL5, nil
	case "LOCAL6":
		return syslog.LOG_LOCAL6, nil
	case "LOCAL7":
		return syslog.LOG_LOCAL7, nil
	default:
		return 0, fmt.Errorf("invalid syslog facility: %s", facility)
	}
}
