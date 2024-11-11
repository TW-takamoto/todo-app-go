package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup() (*zap.Logger, func(), error) {
	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:         "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	logger, err := zapConfig.Build()
	return logger, func() {
		logger.Sync()
	}, err
}
