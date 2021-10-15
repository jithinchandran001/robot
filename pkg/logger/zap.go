package logger

import (
	"runtime/debug"
	"time"

	"go.uber.org/zap"
)

// NewProductionZapLogger will return a new production logger backed by zap
func NewProductionZaplogger() (Logger, error) {
	conf := zap.NewProductionConfig()
	conf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	conf.DisableCaller = true
	conf.DisableStacktrace = true
	zapLogger, err := conf.Build(zap.AddCaller(), zap.AddCallerSkip(1))

	return zpLg{
		lg: zapLogger.Sugar(),
	}, err
}

// NewZapLogger will return a new logger backed by the provided zap instance
func NewZapLogger(lg *zap.Logger) Logger {
	return zpLg{
		lg: lg.Sugar(),
	}
}

type zpLg struct {
	lg *zap.SugaredLogger
}

func (l zpLg) Log(keyValues ...interface{}) error {
	l.lg.Infow("", keyValues...)
	return nil
}

func (l zpLg) With(keyValues ...interface{}) (ll Logger) {
	ll = zpLg{
		lg: l.lg.With(keyValues...),
	}
	return
}

func (l zpLg) Debug(msg string, keyValues ...interface{}) {
	l.lg.With("date", time.Now().UTC().Format(time.RFC3339)).Debugw(msg, keyValues...)
}

func (l zpLg) Info(msg string, keyValues ...interface{}) {
	l.lg.With("date", time.Now().UTC().Format(time.RFC3339)).Infow(msg, keyValues...)
}

func (l zpLg) Warn(msg string, keyValues ...interface{}) {
	l.lg.With("date", time.Now().UTC().Format(time.RFC3339)).Warnw(msg, keyValues...)
}

func (l zpLg) Error(msg string, keyValues ...interface{}) {
	l.lg.With("date", time.Now().UTC().Format(time.RFC3339), "stacktrace", string(debug.Stack())).Errorw(msg, keyValues...)
}

func (l zpLg) ErrorWithoutSTT(msg string, keyValues ...interface{}) {
	l.lg.With("date", time.Now().UTC().Format(time.RFC3339)).Errorw(msg, keyValues...)
}

func (l zpLg) Fatal(msg string, keyValues ...interface{}) {
	l.lg.Fatalw(msg, keyValues...)
}
