package zap-config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func LoggerDefault(options ...zap.Option) (*zap.Logger) {
	return LoggerInfo(options...)
}

// human-readable logger suitable for production code, excludes the caller and stack trace fields
func LoggerInfo(options ...zap.Option) (*zap.Logger) {
	return logger(zapcore.InfoLevel, DefaultEncoderConfig(), options...)
}

// logger suitable for logging in production code, but at debug level.
func LoggerDebug(options ...zap.Option) *zap.Logger {
	return logger(zapcore.DebugLevel, DebugEncoderConfig(), options...)
}

func DebugEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig {		
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func DefaultEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig {		
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		MessageKey:     "M",		
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func logger(level zapcore.Level, encoder zapcore.EncoderConfig, options ...zap.Option) (*zap.Logger) {
	zapConfig := zap.Config {		
		EncoderConfig: encoder,
		Encoding: "console",
		ErrorOutputPaths: []string{"stderr"},
		Level: zap.NewAtomicLevelAt(level),
		OutputPaths: []string{"stdout"},		
	}
	logger, err := zapConfig.Build(options...)
	if err != nil {
		log.Fatal(err)
	}
	return logger
}