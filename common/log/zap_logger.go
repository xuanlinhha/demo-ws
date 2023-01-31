package log

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger
type zapLogger struct {
	logger *zap.Logger
}

// App Logger constructor
func NewLogger(mode, encoding, level string, w io.Writer) Logger {
	// encoder config
	var encoderCfg zapcore.EncoderConfig
	if mode == "Development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// encoder
	var encoder zapcore.Encoder
	if encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	// core
	logWriter := zapcore.AddSync(w)
	logLevel := getLoggerLevel(level)
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(2))
	return &zapLogger{logger: logger}
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(ll string) zapcore.Level {
	level, exist := loggerLevelMap[ll]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func (l *zapLogger) Clean() {
	l.logger.Sync()
}

// Logger methods

func (l *zapLogger) Debug(args ...interface{}) {
	l.logger.Sugar().Debug(args...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Sugar().Debugf(template, args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.logger.Sugar().Info(args...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Sugar().Infof(template, args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.logger.Sugar().Warn(args...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Sugar().Warnf(template, args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.logger.Sugar().Error(args...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Sugar().Errorf(template, args...)
}

func (l *zapLogger) DPanic(args ...interface{}) {
	l.logger.Sugar().DPanic(args...)
}

func (l *zapLogger) DPanicf(template string, args ...interface{}) {
	l.logger.Sugar().DPanicf(template, args...)
}

func (l *zapLogger) Panic(args ...interface{}) {
	l.logger.Sugar().Panic(args...)
}

func (l *zapLogger) Panicf(template string, args ...interface{}) {
	l.logger.Sugar().Panicf(template, args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.logger.Sugar().Fatal(args...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Sugar().Fatalf(template, args...)
}
