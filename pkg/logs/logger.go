package logs

var Log Logger

// NewLogger Init initializes logs
func NewLogger(c *Config) Logger {
	return NewEchoLogger(c)
}

// Logger is a contract for the logger
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

// GetLogger return a log
func GetLogger() Logger {
	return Log
}

// Debug logger
func Debug(args ...interface{}) {
	Log.Debug(args...)
}

// Info logger
func Info(args ...interface{}) {
	Log.Info(args...)
}

// Warn logger
func Warn(args ...interface{}) {
	Log.Warn(args...)
}

// Error logger
func Error(args ...interface{}) {
	Log.Error(args...)
}

// Debugf logger
func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

// Infof logger
func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

// Warnf logger
func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

// Errorf logger
func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}
