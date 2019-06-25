package options

import (
	"encoding/json"
	"unsafe"

	"github.com/spf13/pflag"
)

type Option func(*Options)

type Options struct {
	LogLevel    Level  `json:"level" mapstructure:"level"`               //日志记录级别 ok
	LogFormat   string `json:"format" mapstructure:"format"`             //日志类型,普通 或 json
	Filename    string `json:"filename" mapstructure:"filename"`         //日志保存路径 ok
	MaxSize     int    `json:"max-size" mapstructure:"max-size"`         //日志分割的尺寸 MB //ok
	MaxAge      int    `json:"max-age" mapstructure:"max-age"`           //分割日志保存的时间 day
	Stacktrace  Level  `json:"stacktrace" mapstructure:"stacktrace"`     //记录堆栈的级别
	IsStdOut    bool   `json:"is-stdout" mapstructure:"is-stdout"`       //是否标准输出console输出
	ProjectName string `json:"project-name" mapstructure:"project-name"` //项目名称
	Color       bool   `json:"enable-color" mapstructure:"enable-color"`
}

func NewOptions() *Options {
	return &Options{
		LogLevel:    LogLevel,
		LogFormat:   LogFormat,
		Filename:    Filename,
		MaxSize:     MaxSize,
		MaxAge:      MaxAge,
		Stacktrace:  Stacktrace,
		IsStdOut:    IsStdOut,
		ProjectName: ProjectName,
		Color:       Color,
	}
}

func (o *Options) Validate() []error {
	return nil
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.IntVar((*int)(unsafe.Pointer(&o.LogLevel)), "log.level", int(o.LogLevel), "Minimum log output `LEVEL`")
	fs.StringVar(&o.LogFormat, "log.format", o.LogFormat, "Log output `FORMAT`")
	fs.StringVar(&o.Filename, "log.filename", o.Filename, "Log output filename")
	fs.IntVar(&o.MaxSize, "log.max-size", o.MaxSize, "Divided size of the log (MB)")
	fs.IntVar(&o.MaxAge, "log.max-age", o.MaxAge, "Split log save time (day)")
	fs.IntVar((*int)(unsafe.Pointer(&o.Stacktrace)), "log.stacktrace", int(o.Stacktrace), "Split log save time (day)")
	fs.BoolVar(&o.IsStdOut, "log.is-stdout", o.IsStdOut, "Also log to console")
	fs.StringVar(&o.ProjectName, "log.project-name", o.ProjectName, "Split log save time (day)")
	fs.BoolVar(&o.Color, "log.enable-color", o.Color, "Whether to output colored log")
}

func WithLogLevel(loglevel Level) Option {
	return func(o *Options) {
		o.LogLevel = loglevel
	}
}

func WithLogFormat(logformat string) Option {
	return func(o *Options) {
		o.LogFormat = logformat
	}
}

func WithFilename(logpath string) Option {
	return func(o *Options) {
		o.Filename = logpath
	}
}

func WithMaxSize(maxsize int) Option {
	return func(o *Options) {
		o.MaxSize = maxsize
	}
}

func WithMaxAge(maxage int) Option {
	return func(o *Options) {
		o.MaxAge = maxage
	}
}

func WithStacktrace(stacktrace Level) Option {
	return func(o *Options) {
		o.Stacktrace = stacktrace
	}
}

func WithIsStdOut(isstdout bool) Option {
	return func(o *Options) {
		o.IsStdOut = isstdout
	}
}

func WithProjectName(projectname string) Option {
	return func(o *Options) {
		o.ProjectName = projectname
	}
}

func WithColor(color bool) Option {
	return func(o *Options) {
		o.Color = color
	}
}

func (s *Options) String() string {
	data, _ := json.Marshal(s)
	return string(data)
}
