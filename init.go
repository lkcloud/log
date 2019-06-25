package log

import (
	"os"
	"sync"

	"github.com/lkcloud/log/options"
	"github.com/lkcloud/log/zaplog"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//默认
var l Loger = newLogger(nil)
var once sync.Once
var lock = &sync.RWMutex{}

//设置
func SetLogger(opts ...options.Option) {
	once.Do(func() {
		l = newLogger(nil, opts...)
	})
}

func GetLogger() Loger {
	return l
}

func SetLoggerWithOptions(o *options.Options, opts ...options.Option) {
	once.Do(func() {
		l = newLogger(o, opts...)
	})
}

func Reset() {
	lock.Lock()
	defer lock.Unlock()
	l = newLogger(nil)
}

func newLogger(o *options.Options, opts ...options.Option) *zaplog.Log {
	encoderConfig := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "flag",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if o == nil {
		o = options.NewOptions()
	}

	for _, opt := range opts {
		opt(o)
	}

	// when output to local path, with color is forbidden
	if o.Color && (o.Filename == "") && (o.LogFormat == options.PlainFormat) {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var writers = []zapcore.WriteSyncer{}
	osfileout := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize, // megabytes
		MaxBackups: 3,
		MaxAge:     o.MaxAge, // days
		LocalTime:  true,
		Compress:   false,
	})
	if o.IsStdOut {
		writers = append(writers, os.Stdout)
	}
	writers = append(writers, osfileout)
	w := zapcore.NewMultiWriteSyncer(writers...)

	atom := zap.NewAtomicLevel()
	atom.SetLevel(options.ParseLevel(o.LogLevel)) //改变日志级别

	var enc zapcore.Encoder
	if o.LogFormat == options.PlainFormat {
		enc = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		enc = zapcore.NewJSONEncoder(encoderConfig)
	}
	core := zapcore.NewCore(
		//这里控制json 或者不是json 类型
		enc,
		w,
		atom,
	)
	logger := zap.New(
		core,
		zap.AddStacktrace(options.ParseLevel(o.Stacktrace)),
		zap.AddCaller(),
		zap.AddCallerSkip(2))

	if o.ProjectName != "" {
		logger = logger.With(zap.String(options.ProjectKey, o.ProjectName))
	}
	loggerSugar := logger.Sugar()
	return zaplog.NewLogger(loggerSugar, atom, o)

}

//快捷使用,开发使用
func NewDevelopment(projectName, filePath string) {
	SetLogger(options.WithProjectName(projectName),
		options.WithFilename(filePath),
		options.WithLogFormat(options.JsonFormat),
		options.WithIsStdOut(true))
}

//快捷使用,生产使用
func NewProduction(projectName, filePath string) {
	SetLogger(options.WithProjectName(projectName),
		options.WithLogFormat(options.JsonFormat),
		options.WithFilename(filePath),
		options.WithLogLevel(options.ErrorLevel),
		options.WithIsStdOut(false))
}

//目前只有zap生效
func SetLogLevel(level options.Level) {
	l.SetLogLevel(level)
}

func GetOptions() string {
	return l.Options()
}

//目前只有zap生效
func Sync() {
	l.Sync()
}

//key value
func Debug(msg string, keysAndValues ...interface{}) {
	l.Debug(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	l.Info(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	l.Warn(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	l.Error(msg, keysAndValues...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	l.Panic(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	l.Fatal(msg, keysAndValues...)
}
func Dump(msg string, keysAndValues ...interface{}) {
	l.Dump(msg, keysAndValues...)
}

func Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}
func Dumpf(format string, args ...interface{}) {
	l.Dumpf(format, args...)
}
