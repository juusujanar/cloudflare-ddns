package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.Out = os.Stdout
	log.Level = logrus.InfoLevel
	log.Formatter = &logrus.TextFormatter{FullTimestamp: true}
}


func Info(args ...interface{}) {
	log.Info(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}