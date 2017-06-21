package logger

import (
	"github.com/sirupsen/logrus"
)

type Newlogger struct{
	*logrus.Logger
}

func (l *Newlogger) Critical(args ...interface{}){
	l.Error(args ...)
}

func (l *Newlogger) Criticalf(format string, args ...interface{}){
	l.Errorf(format, args ...)
}

func (l *Newlogger) Notice(args ...interface{}){
	l.Info(args ...)
}

func (l *Newlogger) Noticef(format string, args ...interface{}){
	l.Infof(format, args ...)
}


//TODO: Create test file for the logger^^