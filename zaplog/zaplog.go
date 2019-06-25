package zaplog

import (
	"encoding/json"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/lkcloud/log/options"
	"go.uber.org/zap"
)

type Log struct {
	logger  *zap.SugaredLogger
	atom    zap.AtomicLevel
	options *options.Options
}

func NewLogger(logger *zap.SugaredLogger, atom zap.AtomicLevel, options *options.Options) *Log {
	return &Log{logger: logger, atom: atom, options: options}
}

func (l *Log) Options() string {
	o, _ := json.Marshal(l.options)
	return string(o)
}

//拼接完整的数组
func CoupArray(kv []interface{}) []interface{} {
	if len(kv)%2 != 0 {
		kv = append(kv, kv[len(kv)-1])
		kv[len(kv)-2] = "default"
	}
	return kv
}

func (l *Log) Sync() {
	l.logger.Sync()
}

func (l *Log) SetLogLevel(level options.Level) {
	l.atom.SetLevel(options.ParseLevel(level))
}

func (l *Log) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, CoupArray(keysAndValues)...)
}
func (l *Log) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, CoupArray(keysAndValues)...)
}
func (l *Log) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, CoupArray(keysAndValues)...)
}
func (l *Log) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, CoupArray(keysAndValues)...)
}
func (l *Log) Panic(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, CoupArray(keysAndValues)...)
}
func (l *Log) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, CoupArray(keysAndValues)...)
}
func (l *Log) Dump(msg string, keysAndValues ...interface{}) {
	arr := CoupArray(keysAndValues)
	for k, v := range arr {
		if k%2 == 0 {
			arr[k] = v
		} else {
			arr[k] = strings.Replace(spew.Sdump(v), "\n", "", -1)
		}
	}
	l.logger.Debugw(msg, arr...)
}

func (l *Log) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}
func (l *Log) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}
func (l *Log) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}
func (l *Log) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}
func (l *Log) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}
func (l *Log) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
func (l *Log) Dumpf(format string, args ...interface{}) {
	arrs := make([]interface{}, 0)
	for _, v := range args {
		arrs = append(arrs, strings.Replace(spew.Sdump(v), "\n", "", -1))
	}
	l.logger.Debugf(format, arrs...)
}
