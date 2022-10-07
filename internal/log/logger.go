package log

import (
	"context"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	global *zap.Logger

	globalLevel = zap.DebugLevel

	// onceInit guarantee initialize logger only once
	onceInit sync.Once
)

type Config struct {
	Development bool   `json:"development" mapstructure:"development"`
	Level       string `json:"level" mapstructure:"level"`
}

// Init initializes log by input parameters
// lvl - global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
// timeFormat - custom time format for logger of empty string to use default
func Init() (err error) {
	onceInit.Do(func() {
		// // First, define our level-handling logic.
		// if err = globalLevel.Set(conf.Level); err != nil {
		// 	return
		// }

		// // format log in local as plain text
		// if conf.Development {
		// 	config := zap.NewDevelopmentConfig()
		// 	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// 	config.Level = zap.NewAtomicLevelAt(globalLevel)
		// 	global, err = config.Build(zap.AddStacktrace(zap.ErrorLevel))
		// 	return
		// }

		// High-priority output should also go to standard error, and low-priority
		// output should also go to standard out.
		// It is useful for Kubernetes deployment.
		// Kubernetes interprets os.Stdout log items as INFO and os.Stderr log items
		// as ERROR by default.
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= globalLevel && lvl < zapcore.ErrorLevel
		})
		consoleInfos := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)

		// Configure console output.
		jsonEncoder := newJSONEncoder()
		// Join the outputs, encoders, and level-handling functions into
		// zapcore.
		core := zapcore.NewTee(
			zapcore.NewCore(jsonEncoder, consoleErrors, highPriority),
			zapcore.NewCore(jsonEncoder, consoleInfos, lowPriority),
		)

		// From a zapcore.Core, it's easy to construct a Logger.
		global = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
		zap.RedirectStdLog(global)
	})
	return
}

// Create a new JSON log encoder with the correct settings.
func newJSONEncoder() (encoder zapcore.Encoder) {
	encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	return
}

// Global returns logger with new request_id
func Global() *zap.Logger {
	return global
}

type contextLogger struct {
}

// Logger returns logger which associated context
func Logger(ctx context.Context) *zap.Logger {
	if ctx != nil {
		if v := ctx.Value(contextLogger{}); v != nil {
			return v.(*zap.Logger)
		}
	}
	return Global()
}

// WithLogger returns context which contains logger
func WithLogger(ctx context.Context, fields ...zap.Field) context.Context {
	logger := Logger(ctx)
	if len(fields) > 0 {
		logger = logger.With(fields...)
	}
	return context.WithValue(ctx, contextLogger{}, logger)
}

func Error(ctx context.Context, err error) error {
	Logger(ctx).Error(err.Error())
	return err
}
