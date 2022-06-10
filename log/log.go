package log

import (
	"context"
	"sync"
)

var DefaultLogger Logger
var mu sync.RWMutex

const ContextKeyLogger = "WindealContextKeyLogger"

// InitLogger 初始化 log 组件
func InitLogger(c *Config) {
	DefaultLogger = NewZapLogger(c)
}

// GetDefaultLogger 返回默认的 Logger
func GetDefaultLogger() Logger {
	mu.RLock()
	l := DefaultLogger
	mu.RUnlock()
	return l
}

// SetLevel 设置不同的输出对应的日志级别, output为输出数组下标 "0" "1" "2"
func SetLevel(output string, level Level) {
	GetDefaultLogger().SetLevel(output, level)
}

// GetLevel 获取不同输出对应的日志级别
func GetLevel(output string) Level {
	return GetDefaultLogger().GetLevel(output)
}

// With 日志增加自定义字段，支持多种类型的 Value
func With(fields ...Field) Logger {
	return GetDefaultLogger().With(fields...)
}

// WithContext 以当前 context logger 为基础，增加自定义字段，支持多种类型的 Value
func WithContext(ctx context.Context, fields ...Field) Logger {
	val := ctx.Value(ContextKeyLogger)
	l, ok := val.(Logger)
	if !ok || l == nil {
		l = GetDefaultLogger()
	}

	l = l.With(fields...)
	return l
}

// ContextCloneWith 复制一个新的Context， 新Context会通过key：ContextKeyLogger存储包含带有自定义字段的新的Logger
func ContextCloneWith(ctx context.Context, fields ...Field) context.Context {
	l := WithContext(ctx, fields...)
	newCtx := context.WithValue(ctx, ContextKeyLogger, l)
	return newCtx
}

// GetContextLogger 从 ctx 中获取其绑定的 Logger
func GetContextLogger(ctx context.Context) Logger {
	val := ctx.Value(ContextKeyLogger)
	l, ok := val.(Logger)
	if !ok || l == nil {
		return GetDefaultLogger()
	}
	return l
}

// Debug  打印 Debug 日志，基于 fmt.Print 规则
func Debug(args ...interface{}) {
	GetDefaultLogger().Debug(args...)
}

// Debugf  打印 Debug 日志，基于 fmt.Printf 规则
func Debugf(format string, args ...interface{}) {
	GetDefaultLogger().Debugf(format, args...)
}

// Info  打印 Info 日志，基于 fmt.Print 规则
func Info(args ...interface{}) {
	GetDefaultLogger().Info(args...)
}

// Infof  打印 Info 日志，基于 fmt.Printf 规则
func Infof(format string, args ...interface{}) {
	GetDefaultLogger().Infof(format, args...)
}

// Warn  打印 Warn 日志，基于 fmt.Print 规则
func Warn(args ...interface{}) {
	GetDefaultLogger().Warn(args...)
}

// Warnf  打印 Warn 日志，基于 fmt.Printf 规则
func Warnf(format string, args ...interface{}) {
	GetDefaultLogger().Warnf(format, args...)
}

// Error  打印 Error 日志，基于 fmt.Print 规则
func Error(args ...interface{}) {
	GetDefaultLogger().Error(args...)
}

// Errorf  打印 Error 日志，基于 fmt.Printf 规则
func Errorf(format string, args ...interface{}) {
	GetDefaultLogger().Errorf(format, args...)
}

// Fatal  打印 Fatal 日志，基于 fmt.Print 规则, 会引发 os.Exit
func Fatal(args ...interface{}) {
	GetDefaultLogger().Fatal(args...)
}

// Fatalf  打印 Fatal 日志，基于 fmt.Printf 规则, 会引发 os.Exit
func Fatalf(format string, args ...interface{}) {
	GetDefaultLogger().Fatalf(format, args...)
}

// DebugContext  打印 Debug 日志，基于 fmt.Print 规则
func DebugContext(ctx context.Context, args ...interface{}) {
	GetContextLogger(ctx).Debug(args...)
}

// DebugContextf 打印 Debug 日志, 基于 fmt.Printf 规则
func DebugContextf(ctx context.Context, format string, args ...interface{}) {
	GetContextLogger(ctx).Debugf(format, args...)
}

// InfoContext  打印 Info 日志，基于 fmt.Print 规则
func InfoContext(ctx context.Context, args ...interface{}) {
	GetContextLogger(ctx).Info(args...)
}

// InfoContextf  打印 Info 日志，基于 fmt.Printf 规则
func InfoContextf(ctx context.Context, format string, args ...interface{}) {
	GetContextLogger(ctx).Infof(format, args...)
}

// WarnContext  打印 Warn 日志，基于 fmt.Print 规则
func WarnContext(ctx context.Context, args ...interface{}) {
	GetContextLogger(ctx).Warn(args...)
}

// WarnContextf  打印 Warn 日志，基于 fmt.Printf 规则
func WarnContextf(ctx context.Context, format string, args ...interface{}) {
	GetContextLogger(ctx).Warnf(format, args...)
}

// ErrorContext  打印 Error 日志，基于 fmt.Print 规则
func ErrorContext(ctx context.Context, args ...interface{}) {
	GetContextLogger(ctx).Error(args...)
}

// ErrorContextf  打印 Error 日志，基于 fmt.Printf 规则
func ErrorContextf(ctx context.Context, format string, args ...interface{}) {
	GetContextLogger(ctx).Errorf(format, args...)
}

// FatalContext  打印 Fatal 日志，基于 fmt.Print 规则, 会引发 os.Exit
func FatalContext(ctx context.Context, args ...interface{}) {
	GetContextLogger(ctx).Fatal(args)
}

// FatalContextf  打印 Fatal 日志，基于 fmt.Printf 规则, 会引发 os.Exit
func FatalContextf(ctx context.Context, format string, args ...interface{}) {
	GetContextLogger(ctx).Fatalf(format, args...)
}
