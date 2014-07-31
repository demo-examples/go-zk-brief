package main

import (
	"net/http"
	"flag"
)

func main() {
	flag.Parse()
	logInfof("spr_api version:%s started...\n", VERSION)

	if err := InitConfig(); err != nil {
		logError("error when init config...")
		panic(err)
	}
	if err := initZK(); err != nil {
		logError("error when init zk...")
		panic(err)
	}

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
		logInfo("ListenAndServe:", err)
	}

}

