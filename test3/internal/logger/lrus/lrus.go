package lrus

import (
	"github.com/sirupsen/logrus"
	"io"
)

type Lrus struct {
	Log *logrus.Logger
}

func New(level string, output io.Writer) *Lrus {
	l := &Lrus{Log: logrus.New()}
	l.Log.Formatter = &logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
	}
	l.Log.Level = l.SetLevel(level)
	l.Log.SetOutput(output)
	return l
}

func (l *Lrus) SetLevel(level string) (lev logrus.Level) {
	switch level {
	case "INFO":
		lev = logrus.InfoLevel
	case "ERROR":
		lev = logrus.ErrorLevel
	case "DEBUG":
		lev = logrus.DebugLevel
	case "WARN":
		lev = logrus.WarnLevel
	default:
		lev = logrus.InfoLevel
	}
	return
}

/*func (l *Lrus) SetOutput(output io.Writer) {
	l.Log.SetOutput(output)
}*/

func (l *Lrus) Info(msg interface{}) {
	l.Log.Infoln(msg)
}

func (l *Lrus) Error(msg interface{}) {
	l.Log.Error(msg)
}

func (l *Lrus) Debug(msg interface{}) {
	l.Log.Debug(msg)
}

func (l *Lrus) Warn(msg interface{}) {
	l.Log.Warn(msg)
}
