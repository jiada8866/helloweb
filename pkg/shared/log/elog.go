package log

import (
	"encoding/json"
	"io"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type ELog struct {
	prefix string

	*logrus.Logger
}

func NewELog(prefix string) *ELog {
	return &ELog{
		prefix: prefix,
		Logger: logrus.StandardLogger(),
	}
}

func (el *ELog) Output() io.Writer {
	return el.Out
}

func (el *ELog) SetOutput(w io.Writer) {
	el.Out = w
}

func (el *ELog) Prefix() string {
	return el.prefix
}

func (el *ELog) SetPrefix(p string) {
	el.prefix = p
}

func (el *ELog) Level() log.Lvl {
	return log.Lvl(el.Logger.Level)
}

func (el *ELog) SetLevel(l log.Lvl) {
	var level = logrus.InfoLevel
	switch l {
	case log.DEBUG:
		level = logrus.DebugLevel
	case log.INFO:
		level = logrus.InfoLevel
	case log.WARN:
		level = logrus.WarnLevel
	case log.ERROR:
		level = logrus.ErrorLevel
	case log.OFF:
		level = logrus.FatalLevel
	}
	el.Logger.Level = level
}

func (el *ELog) Printj(j log.JSON) {

}

func (el *ELog) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		el.Fatal(err)
	}
	el.Debug(string(b))
}

func (el *ELog) Infoj(j log.JSON) {

}

func (el *ELog) Warnj(j log.JSON) {

}

func (el *ELog) Errorj(j log.JSON) {

}

func (el *ELog) Fatalj(j log.JSON) {

}

func (el *ELog) Panicj(j log.JSON) {

}
