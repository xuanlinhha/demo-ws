package log

import "os"

var (
	defaultLogger = NewLogger(
		"Development",
		"console",
		"debug",
		os.Stdout,
	)
)

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}
func DPanic(args ...interface{}) {
	defaultLogger.DPanic(args...)
}
func DPanicf(format string, args ...interface{}) {
	defaultLogger.DPanicf(format, args...)
}
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}
