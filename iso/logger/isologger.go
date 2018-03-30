package logger

import (
	"log"
	"strings"
)

type logFlags struct {
	trace   bool
	info    bool
	warning bool
	error   bool
}

type loggerBase struct {
}

func (l *loggerBase) parse(lF uint8) *logFlags {

	return &logFlags{
		trace:   lF&Trace == Trace,
		info:    lF&Info == Info,
		warning: lF&Warning == Warning,
		error:   lF&Error == Error,
	}

}

type ISOLogger struct {
	loggerBase
	stdout *log.Logger
	stderr *log.Logger
}

func NewLogger(sout *log.Logger, serr *log.Logger) *ISOLogger {
	return &ISOLogger{
		stdout: sout,
		stderr: serr,
	}
}

func (l *ISOLogger) Log(lf uint8, format string, v ...interface{}) {

	flg := l.parse(lf)
	format = strings.TrimSpace(format)
	if format[len(format)-1] != '\n' {
		format += "\n"
	}
	if flg.trace || flg.info || flg.warning {
		l.stdout.Printf(format, v)
	}
	if flg.error {
		l.stderr.Printf(format, v)
	}
}
