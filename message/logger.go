package message

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func LogContext(ctx string, cor ...interface{}) *logrus.Entry {
	entry := logrus.WithFields(logrus.Fields{
		"topic":   "go-chankey",
		"context": ctx,
		"at":      time.Now().Format("2006-01-02 15:04:05"),
	})

	if len(cor) > 0 {
		if cor[0] != nil {
			entry = entry.WithFields(
				logrus.Fields{
					"correlation-id": fmt.Sprintf("%+v", cor[0]),
				})
		}
	}

	return entry
}

func Log(level logrus.Level, message, context string, corr ...interface{}) {

	var correlation interface{}
	if len(corr) > 0 {
		correlation = corr[0]
	}

	entry := LogContext(context, correlation)
	switch level {
	case logrus.DebugLevel:
		entry.Debug(message)
	case logrus.InfoLevel:
		entry.Info(message)
	case logrus.WarnLevel:
		entry.Warn(message)
	case logrus.ErrorLevel:
		entry.Error(message)
	}
}

func GetLogger() *logrus.Logger {
	return logrus.StandardLogger()
}
