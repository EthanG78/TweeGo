package logger_test

import (
	"testing"
	"github.com/EthanG78/tweego/logger"
	"github.com/sirupsen/logrus"
)



var log = &logger.Newlogger{Logger: logrus.New()}

func TestNewlogger_Critical(t *testing.T) {
	log.Critical("This is a test")
}

func TestNewlogger_Criticalf(t *testing.T) {
	log.Criticalf("This is a test, logger is of type %T", log)
}

func TestNewlogger_Notice(t *testing.T) {
	log.Notice("This is a test")
}

func TestNewlogger_Noticef(t *testing.T) {
	log.Noticef("This is a test, logger is of type %T", log)
}