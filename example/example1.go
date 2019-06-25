package main

import (
	"github.com/lkcloud/log"
	"github.com/lkcloud/log/options"
	"github.com/spf13/pflag"
)

var help bool

func main() {
	logOption := options.NewOptions()
	logOption.AddFlags(pflag.CommandLine)
	pflag.CommandLine.BoolVarP(&help, "help", "h", false, "help flag")
	pflag.Parse()

	log.SetLoggerWithOptions(logOption, options.WithIsStdOut(true))
	defer log.Sync()

	if help {
		pflag.Usage()
	}

	log.Infof("log options is: `%s`", log.GetOptions())

	log.Info("this is a test log")
	log.Error("error: this is a test log")
	log.Warn("this is a test log", "lable", "value")
	log.Infof("infof: %s", "args")

}
