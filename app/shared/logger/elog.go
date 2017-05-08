package logger

import (
	"encoding/json"
	"io"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/gommon/log"
)

type Elog struct {
	prefix string

	*logrus.Logger
}

func (el *Elog) Output() io.Writer {
	return el.Out
}

func (el *Elog) SetOutput(w io.Writer) {
	el.Out = w
}

func (el *Elog) Prefix() string {
	return el.prefix
}

func (el *Elog) SetPrefix(p string) {
	el.prefix = p
}

func (el *Elog) Level() log.Lvl {
	return log.Lvl(el.Logger.Level)
}

func (el *Elog) SetLevel(l log.Lvl) {
	var level logrus.Level = logrus.InfoLevel
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

func (el *Elog) Printj(j log.JSON) {

}

func (el *Elog) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		el.Fatal(err)
	}
	el.Debug(string(b))
}

func (el *Elog) Infoj(j log.JSON) {

}

func (el *Elog) Warnj(j log.JSON) {

}

func (el *Elog) Errorj(j log.JSON) {

}

func (el *Elog) Fatalj(j log.JSON) {

}

func (el *Elog) Panicj(j log.JSON) {

}
