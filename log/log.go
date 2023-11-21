package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// TODO: add cli config file to provide the ability to change log level and log file path
const (
	logFile = "/tmp/ff.log"
)

var logger *logrus.Logger

type Fields map[string]any

func WithFields(fields Fields) *logrus.Entry {
	return logger.WithFields(logrus.Fields(fields))
}

func Info(args ...any) {
	logger.Info(args...)
}

func Warn(args ...any) {
	logger.Warn(args...)
}

func Error(args ...any) {
	logger.Error(args...)
}

func Fatal(args ...any) {
	logger.Fatal(args...)
}

func Infof(format string, args ...any) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...any) {
	logger.Warnf(format, args...)
}
func Debugf(format string, args ...any) {
	logger.Debugf(format, args...)
}

func Errorf(format string, args ...any) {
	logger.Errorf(format, args...)
}

func ClearLogFile() error {
	logFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	logFile.Truncate(0)
	logFile.Close()
	return nil
}

func init() {
	logger = logrus.New()
	setupLogger(logger)
}

func setupLogger(l *logrus.Logger) {
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)

	l.Out = os.Stdout
	logFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		l.Out = logFile
	} else {
		l.Info("Failed to log to file, using default stderr")
	}
}
