package logger

import (
	"fmt"
	"github.com/armanokka/transloapi/internal/usecase"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	log *zap.Logger
}

func (l Logger) Info(message string, args ...interface{}) {
	log := l.with(args...)
	log.log.Info(message)
}

func (l Logger) Debug(message string, args ...interface{}) {
	log := l.with(args...)
	log.log.Debug(message)
}

func (l Logger) Warn(message string, args ...interface{}) {
	log := l.with(args...)
	log.log.Warn(message)
}

func (l Logger) Error(message string, args ...interface{}) {
	log := l.with(args...)
	log.log.Error(message)
}

func (l Logger) Fatal(message string, args ...interface{}) {
	log := l.with(args...)
	log.log.Fatal(message)
	os.Exit(1)
}

func (l Logger) with(args ...interface{}) Logger {
	log := l.log.With()
	for _, arg := range args {
		switch v := arg.(type) {
		case zapcore.Field:
			log = log.With(v)
		case []interface{}:
			return l.with(v)
		default:
			log.Error("with: unknown type", zap.String("type", fmt.Sprint("%T", arg))) //nolint:govet
		}
	}
	return Logger{log: log}
}

func (l Logger) With(args ...interface{}) usecase.Logger {
	log := l.log.With()
	for _, arg := range args {
		switch v := arg.(type) {
		case zapcore.Field:
			log = log.With(v)
		case []interface{}:
			return l.With(v)
		default:
			log.Error("with: unknown type", zap.String("type", fmt.Sprintf("%T", arg))) //nolint:govet
		}
	}
	return Logger{log: log}
}

func New(level string) Logger {
	l, err := zapcore.ParseLevel(level)
	if err != nil {
		l = zapcore.DebugLevel
	}
	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
	log := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.AddSync(colorable.NewColorableStdout()),
		l,
	), zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	defer func() {
		err = log.Sync()
		if err != nil {
			fmt.Println("log sync:", err)
		}
	}()
	return Logger{log: log}
}
