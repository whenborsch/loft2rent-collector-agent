package logger

import (
	"github.com/prometheus/client_golang/prometheus"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//nolint:gochecknoglobals,gomnd
var (
	errorsCounter *prometheus.CounterVec
)

type Logger struct {
	logger *zap.Logger
}

var globalLog *Logger

func Init(logLevel, format string) (err error) {
	globalLog, err = build(logLevel, format)

	return err
}

func GetInstance() *Logger {
	if globalLog == nil {
		_ = Init("info", "console")
	}

	return globalLog
}

func build(logLevel, format string) (*Logger, error) {
	level := zap.NewAtomicLevel()
	switch logLevel {
	case "debug":
		level.SetLevel(zapcore.DebugLevel)
	case "info":
		level.SetLevel(zapcore.InfoLevel)
	case "warn":
		level.SetLevel(zapcore.WarnLevel)
	case "error":
		level.SetLevel(zapcore.ErrorLevel)
	default:
		level.SetLevel(zapcore.InfoLevel)
	}

	if format != "json" {
		format = "console"
	}

	WorkErrorsInitialize()

	logger, err := zap.Config{
		Level:    level,
		Encoding: format,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:  "time",
			LevelKey: "level",
			NameKey:  "logger",
			// CallerKey: "Source",
			// StacktraceKey: "St",
			MessageKey:     "action",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()

	if err != nil {
		return nil, err
	}

	return &Logger{logger: logger}, nil
}

func (log *Logger) Fatal(event string, fields ...zap.Field) {
	log.logger.Fatal(event, fields...)
	errorsCounter.WithLabelValues(event, "fatal").Inc()
}

func (log *Logger) Error(event string, fields ...zap.Field) {
	log.logger.Error(event, fields...)
	errorsCounter.WithLabelValues(event, "error").Inc()
}

func (log *Logger) Warning(event string, fields ...zap.Field) {
	log.logger.Warn(event, fields...)
	errorsCounter.WithLabelValues(event, "warning").Inc()
}

func (log *Logger) Info(event string, fields ...zap.Field) {
	log.logger.Info(event, fields...)
}

func (log *Logger) Debug(event string, fields ...zap.Field) {
	log.logger.Debug(event, fields...)
}

func WorkErrorsInitialize() {
	errorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "errors_count",
			Help: "Errors counter by logger",
		},
		[]string{"method", "level"})

	prometheus.MustRegister(errorsCounter)
}
