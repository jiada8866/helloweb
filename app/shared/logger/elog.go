package logger

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/gommon/log"
	"io"
)

type Elog struct {
	*logrus.Logger
}

func (e *Elog) SetOutput(w io.Writer) {
	e.Out = w
}

func (e *Elog) SetLevel(l log.Lvl) {
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
	e.Level = level
}

func (e *Elog) Printj(j log.JSON) {

}

func (e *Elog) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		e.Fatal(err)
	}
	e.Debug(string(b))
}

func (e *Elog) Infoj(j log.JSON) {

}

func (e *Elog) Warnj(j log.JSON) {

}

func (e *Elog) Errorj(j log.JSON) {

}

func (e *Elog) Fatalj(j log.JSON) {

}
