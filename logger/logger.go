package logger

import "go.uber.org/zap"

type Logger interface {
	Notice(msg string, field ...zap.Field)
	Info(msg string, field ...zap.Field)
	Warn(msg string, field ...zap.Field)
	Error(msg string, field ...zap.Field)
	Debug(msg string, field ...zap.Field)
	PublishLog(logData map[string]interface{})
}
