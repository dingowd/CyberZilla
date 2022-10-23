package standart

import (
	"io"
	"log"
)

type Logger struct {
	warn  *log.Logger
	info  *log.Logger
	err   *log.Logger
	debug *log.Logger
}

func New(level string, output io.Writer) *Logger {
	l := &Logger{}
	l.warn = log.New(output, "WARNING: ", log.Ldate|log.Ltime)
	l.info = log.New(output, "INFO: ", log.Ldate|log.Ltime)
	l.err = log.New(output, "ERROR: ", log.Ldate|log.Ltime)
	l.debug = log.New(output, "DEBUG: ", log.Ldate|log.Ltime)

	return l
}

func (l *Logger) Info(msg interface{}) {
	l.info.Println(msg)
}

func (l *Logger) Error(msg interface{}) {
	l.err.Println(msg)
}

func (l *Logger) Debug(msg interface{}) {
	l.debug.Println(msg)
}

func (l *Logger) Warn(msg interface{}) {
	l.warn.Println(msg)
}
