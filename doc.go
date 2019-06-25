/*
Package log is a structured logger for Go, completely API compatible with the standard library logger.

The simplest way to use log is simply the package-level exported logger:

package main

import "github.com/lkcloud/log"

func main() {
	log.Info("this is a test message", "level", "info")
}

Output:
  2019-06-25 11:43:11.170	INFO	example/simple.go:6	this is a test message	{"level": "info"}

For a full guide visit https://github.com/lkcloud/log
*/
package log // import "github.com/lkcloud/log"
