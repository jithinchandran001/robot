package logger

// Logger interface that will abstract logging functionality
type Logger interface {
	Log(keyvals ...interface{}) error
	With(keyvals ...interface{}) Logger

	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Warn(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
	ErrorWithoutSTT(msg string, keyvals ...interface{})
	Fatal(msg string, keyvals ...interface{})
}
