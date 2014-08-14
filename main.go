package main

import (
	"net/http"
	"flag"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Infof("spr_api version:%s started...\n", VERSION)

	if err := InitConfig(); err != nil {
		glog.Error("error when init config...")
		panic(err)
	}
	if err := initZK(); err != nil {
		glog.Error("error when init zk...")
		panic(err)
	}

	go backup()   // 定时备份zk

	http.HandleFunc("/servicelist", servicelist)
	http.HandleFunc("/serverlist", serverlist)
//	http.HandleFunc("/createnode", createnode)
	http.HandleFunc("/addservice", addservice)
	http.HandleFunc("/delservice", delservice)
	http.HandleFunc("/addserver", addserver)
	http.HandleFunc("/delserver", delserver)

	debugf("[%T]%v\n", Conf.HttpBind, Conf.HttpBind)

	err := http.ListenAndServe(Conf.HttpBind, nil)
	if err != nil {
		glog.Errorf("ListenAndServe:", err)
	}

}

