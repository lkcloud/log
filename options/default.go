package options

import (
	"go.uber.org/zap/zapcore"
)

//日志级别
type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	}
	return "unknown"
}

func ParseLevel(level Level) zapcore.Level {
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case PanicLevel:
		return zapcore.PanicLevel
	case FatalLevel:
		return zapcore.FatalLevel
	}

	return zapcore.DebugLevel
}

var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
}

// log format
const (
	JsonFormat  = "json"
	PlainFormat = "plain"
)

//默认参数
const (
	LogLevel    Level  = DebugLevel //日志记录级别
	LogFormat   string = PlainFormat
	Filename    string = ""         //日志保存路径 //需要设置程序当前运行路径
	MaxSize     int    = 100        //日志分割的尺寸 MB
	MaxAge      int    = 30         //分割日志保存的时间 day
	Stacktrace  Level  = PanicLevel //记录堆栈的级别
	IsStdOut    bool   = true       //是否标准输出console输出
	ProjectName string = ""         //项目名称
	ProjectKey  string = "service"  //
	Color       bool   = true
	ConfigFile  string = ""
)
