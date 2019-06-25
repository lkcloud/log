package log

import (
	"github.com/lkcloud/log/options"
)

//定义接口
type Loger interface {
	//key value
	Debug(string, ...interface{}) //调试的
	Info(string, ...interface{})  //提示的
	Warn(string, ...interface{})  //警告的
	Error(string, ...interface{}) //错误的
	Panic(string, ...interface{}) //恐慌的
	Fatal(string, ...interface{}) //致命的
	Dump(string, ...interface{})  //详细结构类型,调试利器
	// format output
	Debugf(string, ...interface{}) //调试的
	Infof(string, ...interface{})  //提示的
	Warnf(string, ...interface{})  //警告的
	Errorf(string, ...interface{}) //错误的
	Panicf(string, ...interface{}) //恐慌的
	Fatalf(string, ...interface{}) //致命的
	Dumpf(string, ...interface{})  //详细结构类型,调试利器
	Sync()                         //同步
	Options() string
	SetLogLevel(options.Level) //可以随机设置日志级别的
}
