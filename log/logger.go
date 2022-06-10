package log

import (
	"strings"
)

// Level  日志等级类型
type Level int

// Field 日志自定义字段
type Field struct {
	Key   string
	Value interface{}
}

// 枚举日志等级
const (
	LevelNil Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// LevelToStr 将日志等级转化为字符串
func LevelToStr(lv *Level) string {
	switch *lv {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	default:
		return ""
	}
}

// StrToLevel 字符串表示的日志等级转化为Level
func StrToLevel(level string) Level {
	switch strings.ToLower(level) {
	case "trace":
		return LevelTrace
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warn":
		return LevelWarn
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	default:
		return LevelNil
	}
}

// Logger 是对外提供的日志接口
type Logger interface {
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{}) // Fatal日志会引发os.Exit(1)
	Fatalf(format string, args ...interface{})

	// Sync 清洗缓存， 在程序结束前应该调用Sync将缓冲区的日志写到目标位置
	Sync() error

	// SetLevel 设置输出端日志级别
	SetLevel(output string, level Level) error
	// GetLevel 获取输出端日志级别
	GetLevel(output string) Level

	// With 日志增加自定义字段，
	With(fields ...Field) Logger
}
