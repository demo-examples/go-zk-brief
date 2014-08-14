package main

import(
//	"fmt"
	"github.com/golang/glog"
)
func init() {
//	fmt.Println("log start...")
	defer glog.Flush()
}

func apilog(input string, api string, rtn string) {
	glog.Infof("[%s]params:%s\t rtn:%s", api, input, rtn)
}

func debug(args ...interface{}) {
//	fmt.Println(args)
}

func debugf(format string, args ...interface{}) {
//	fmt.Printf(format, args...)
}
