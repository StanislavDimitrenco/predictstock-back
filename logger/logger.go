package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// Logger enforces specific log message formats
type Logger struct {
	logrus *logrus.Logger
	fields logrus.Fields
}

type Fields map[string]interface{}

// NewLogger initializes the standard logger
func NewLogger(fileName string) *Logger {
	var baseLogger = logrus.New()

	var logger = &Logger{baseLogger, logrus.Fields{}}

	logger.logrus.Formatter = &logrus.TextFormatter{
		DisableSorting:  false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	file, err := os.OpenFile(fmt.Sprintf("logs/%s.log", fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		logger.logrus.Out = io.Writer(file)
	} else {
		logger.logrus.Info("Failed to log to file, using default stderr")
	}

	return logger
}

func (l *Logger) SetFields(arguments map[string]interface{}) *Logger {
	for k, v := range arguments {
		l.fields[k] = v
	}

	return l
}

func (l *Logger) LogInfo(args ...interface{}) {
	l.logrus.WithFields(l.fields).Infoln(args)
}

func (l *Logger) LogWarning(args ...interface{}) {
	l.logrus.WithFields(l.fields).Warningln(args)
}

func (l *Logger) LogError(args ...interface{}) {
	l.logrus.WithFields(l.fields).Errorln(args)
}
