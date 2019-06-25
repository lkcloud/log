package log

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/lkcloud/log/options"
)

func TestLogger(t *testing.T) {
	SetLogger(options.WithLogFormat(options.JsonFormat), //打印json格式
		options.WithProjectName("k3日志"),          //设置项目名称
		options.WithFilename("log.txt"),          //设置输出文件名,或输出的路径
		options.WithLogLevel(options.ErrorLevel), //设置日志级别,默认debug
		options.WithMaxAge(30),                   //日志保存天数,默认30天
		options.WithMaxSize(512),                 //多少M进行分隔日志,默认100M
		//WithStacktrace(PanicLevel),                   //设置堆栈级别
		options.WithIsStdOut(true)) //是否同时输出控制台
	defer Sync()
	Debug("debug日志", 1)
	Info("info日志", 2)
	Warn("warn日志", 3)
	Error("error日志", 4)
	//Panic("panic", 5)
	//Fatal("fatal", 6)
}

func BenchmarkInfo(t *testing.B) {
	t.ResetTimer()
	runtime.GOMAXPROCS(runtime.NumCPU())
	SetLogger(options.WithLogFormat(options.JsonFormat), //打印json格式
		options.WithProjectName("k3日志"),         //设置项目名称
		options.WithFilename("log.txt"),         //设置输出文件名,或输出的路径
		options.WithLogLevel(options.InfoLevel), //设置日志级别,默认debug
		options.WithMaxAge(30),                  //日志保存天数,默认30天
		options.WithMaxSize(10),                 //多少M进行分隔日志,默认100M
		options.WithStacktrace(2),               //设置堆栈级别
		options.WithIsStdOut(false))             //是否同时输出控制台
	defer Sync()
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		Info("测试日志", "打印结果", 100)
	}
}
func TestDebug(t *testing.T) {
	SetLogger(options.WithIsStdOut(true),
		options.WithLogFormat(options.JsonFormat))
	Debug("ddd", 200, "aa", 2001, "bb")
}
func TestError(t *testing.T) {
	Error("error", 1)
	fmt.Println("继续执行")
}
func TestInfo2(t *testing.T) {
	Info("aa", 11)
	SetLogLevel(options.InfoLevel)
	Info("info", 100)
	Warn("warn", 200)
	SetLogLevel(options.ErrorLevel)
	Info("info-100", 300) //这个无法输出,因为上面设置日志级别为:error
	Error("err", 400)
}
func TestDump(t *testing.T) {
	type s struct {
		Name string
		Age  int
	}
	SetLogger(options.WithIsStdOut(true))
	Dump("name", "dump", "s", s{Name: "k3", Age: 2})
}
