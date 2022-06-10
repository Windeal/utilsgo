// Package log
// @file zaplogger.go: 基于go.uber.org/zap实现的Logger
package log

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"strconv"
)

// zapLog 基于zaplogger的Logger实现
type zapLog struct {
	levels []*zap.AtomicLevel
	logger *zap.Logger
}

func NewZapLogger(c *Config) *zapLog {

	var cores []zapcore.Core
	var levels []*zap.AtomicLevel
	callerSkip := 2

	for _, oc := range c.Outputs {

		switch oc.Writer {
		case "file":
			core, al, err := newFileCore(&oc)
			if err != nil {
				panic("new file core error, " + err.Error())
			}
			cores = append(cores, core)
			levels = append(levels, &al)
		case "console":
			core, al, err := newConsoleCore(&oc)
			if err != nil {
				panic("new console core error, " + err.Error())
			}
			cores = append(cores, core)
			levels = append(levels, &al)
		default:
			core, al, err := newConsoleCore(&oc)
			if err != nil {
				panic("new console core error, " + err.Error())
			}
			cores = append(cores, core)
			levels = append(levels, &al)

		}
	}

	return &zapLog{
		levels: levels,
		logger: zap.New(
			zapcore.NewTee(cores...),
			zap.AddCallerSkip(callerSkip),
			zap.AddCaller(),
		),
	}
}

// With 日志增加自定义字段
func (l *zapLog) With(fields ...Field) Logger {
	zapFields := make([]zap.Field, len(fields))
	for i := range fields {
		zapFields[i] = zap.Any(fields[i].Key, fields[i].Value)
	}

	return &zapLog{logger: l.logger.With(zapFields...)}
}

// Trace :
func (l *zapLog) Trace(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.DebugLevel) {
		msg := fmt.Sprint(args...)
		l.logger.Debug(msg)
	}
}

// Tracef :
func (l *zapLog) Tracef(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.DebugLevel) {
		msg := fmt.Sprintf(format, args...)
		l.logger.Debug(msg)
	}
}

// Debug :
func (l *zapLog) Debug(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.DebugLevel) {
		msg := fmt.Sprint(args...)
		l.logger.Debug(msg)
	}
}

// Debugf :
func (l *zapLog) Debugf(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.DebugLevel) {
		msg := fmt.Sprintf(format, args...)
		l.logger.Debug(msg)
	}
}

// Info :
func (l *zapLog) Info(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.InfoLevel) {
		msg := fmt.Sprint(args...)
		l.logger.Info(msg)
	}
}

// Infof :
func (l *zapLog) Infof(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.InfoLevel) {
		msg := fmt.Sprintf(format, args...)
		l.logger.Info(msg)
	}
}

// Warn :
func (l *zapLog) Warn(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.WarnLevel) {
		msg := fmt.Sprint(args...)
		l.logger.Warn(msg)
	}
}

// Warnf :
func (l *zapLog) Warnf(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.WarnLevel) {
		msg := fmt.Sprintf(format, args...)
		l.logger.Warn(msg)
	}
}

// Error :
func (l *zapLog) Error(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.ErrorLevel) {
		msg := fmt.Sprint(args...)
		l.logger.Error(msg)
	}
}

// Errorf :
func (l *zapLog) Errorf(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.ErrorLevel) {
		msg := fmt.Sprintf(format, args...)
		l.logger.Error(msg)
	}
}

// Fatal :
func (l *zapLog) Fatal(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.FatalLevel) {
		msg := fmt.Sprint(args...)
		l.logger.Fatal(msg)
	}
}

// Fatalf :
func (l *zapLog) Fatalf(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.FatalLevel) {
		msg := fmt.Sprintf(format, args...)
		l.logger.Fatal(msg)
	}
}

// Sync : 冲洗缓冲区的日志，应用程序最好在退出前要调用下，以免有日志在缓冲区没有写入到目标位置。
func (l *zapLog) Sync() error {
	return l.logger.Sync()
}

// SetLevel 设置输出端日志级别
func (l *zapLog) SetLevel(output string, level Level) error {
	i, e := strconv.Atoi(output)
	if e != nil {
		return e
	}
	if i < 0 || i >= len(l.levels) {
		return errors.New("invalid output, out of range")
	}
	l.levels[i].SetLevel(levelToZapLevel(level))
	return nil
}

// GetLevel 获取输出端日志级别
func (l *zapLog) GetLevel(output string) Level {
	i, e := strconv.Atoi(output)
	if e != nil {
		return LevelDebug
	}
	if i < 0 || i >= len(l.levels) {
		return LevelDebug
	}
	return zapLevelToLevel(l.levels[i].Level())
}

func zapLevelToLevel(lv zapcore.Level) Level {
	switch lv {
	case zapcore.DebugLevel:
		return LevelDebug
	case zapcore.InfoLevel:
		return LevelInfo
	case zapcore.WarnLevel:
		return LevelWarn
	case zapcore.ErrorLevel:
		return LevelError
	case zapcore.FatalLevel:
		return LevelFatal
	case zapcore.PanicLevel:
		return LevelFatal
	default:
		return LevelWarn
	}
}

func levelToZapLevel(lv Level) zapcore.Level {
	switch lv {
	case LevelTrace:
		return zapcore.DebugLevel
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zap.WarnLevel
	case LevelError:
		return zap.ErrorLevel
	case LevelFatal:
		return zap.FatalLevel
	default:
		return zap.WarnLevel
	}
}

func newConsoleCore(c *OutputConfig) (zapcore.Core, zap.AtomicLevel, error) {
	lvl := zap.NewAtomicLevelAt(levelToZapLevel(StrToLevel(c.Level)))
	return zapcore.NewCore(
		newEncoder(c),
		zapcore.Lock(os.Stdout),
		lvl), lvl, nil
}

func newFileCore(c *OutputConfig) (zapcore.Core, zap.AtomicLevel, error) {

	logfile := path.Join(c.WriterConfig.LogPath, c.WriterConfig.Filename)

	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    c.WriterConfig.MaxSize,
		MaxBackups: c.WriterConfig.MaxBackups,
		MaxAge:     c.WriterConfig.MaxAge,
	})

	// 日志级别
	lvl := zap.NewAtomicLevelAt(levelToZapLevel(StrToLevel(c.Level)))
	return zapcore.NewCore(
		newEncoder(c),
		ws, lvl,
	), lvl, nil
}

func newEncoder(c *OutputConfig) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	switch c.Formatter {
	case "console":
		return zapcore.NewConsoleEncoder(encoderConfig)
	case "json":
		return zapcore.NewJSONEncoder(encoderConfig)
	default:
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}
