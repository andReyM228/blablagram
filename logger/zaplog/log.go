package zaplog

import (
	"blablagram/logger"
	"go.uber.org/zap"
)

// ensures that zaplog implements logger.Logger.
var _ logger.Logger = (*zaplog)(nil)

// zaplog is an implementation of logger.Logger using zap.
type zaplog struct {
	client *zap.Logger
}

// NewLog is a constructor for a logger.Logger.
func NewLog() logger.Logger {
	return &zaplog{
		client: zap.NewExample(),
	}
}

// Debug is used to send formatted debug message.
func (log zaplog) Debug(msg string) {
	log.client.Debug(msg)
}

// Error is used to send formatted as error message.
func (log zaplog) Error(msg string, err error) {
	log.client.Error(msg, zap.Error(err))
}

// Fatal is used to send formatted as error message and interact program.
func (log zaplog) Fatal(msg string, err error) {
	log.client.Fatal(msg, zap.Error(err))
}
