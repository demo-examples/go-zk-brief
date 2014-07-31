package main

import(
	"fmt"
	"github.com/golang/glog"
)
func init() {
//	fmt.Println("log start...")
	defer glog.Flush()
}
func logInfo(msg ...interface{}) {
	glog.Info(msg...)
}

func logError(msg ...interface{}) {
	glog.Error(msg...)
}

func logFatal(msg ...interface{}) {
	glog.Fatal(msg...)
}

func logInfof(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func logErrorf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}

func logFatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args...)
}

func apilog(input string, api string, rtn string) {
	glog.Infof("[%s]params:%s\t rtn:%s", api, input, rtn)
}



func debug(args ...interface{}) {
	fmt.Println(args)
}

func debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
