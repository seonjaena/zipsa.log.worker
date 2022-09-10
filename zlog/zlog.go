package zlog

import (
	"github.com/google/martian/log"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	zp "zipsa.log.worker/properties"
)

const (
	outStdout string = "stdout"
	outFile   string = "file"
)

var instance = logrus.New()

func init() {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true
	if strings.Compare(zp.GetLogOut(), outStdout) == 0 {
		formatter.ForceColors = true
	}
	instance.SetFormatter(formatter)

	switch zp.GetLogOut() {
	case outStdout:
		instance.SetOutput(os.Stdout)
	case outFile:
		var filename = "logs/logfile.log"

		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			instance.Errorf("zlog os.OpenFile error msg=%s", err)
		} else {
			instance.SetOutput(f)
		}
	default:
		log.Errorf("log out is invalid. logout=%s", zp.GetLogOut())
	}

	switch zp.GetLogLevel() {
	case "debug":
		instance.SetLevel(logrus.DebugLevel)
	case "warn":
		instance.SetLevel(logrus.WarnLevel)
	case "info":
		instance.SetLevel(logrus.InfoLevel)
	case "error":
		instance.SetLevel(logrus.ErrorLevel)
	default:
		log.Errorf("log level is invalid. loglevel=%s", zp.GetLogLevel())
	}

	instance.Info("zlog init complete.")
}

func Instance() *logrus.Logger {
	return instance
}
