package main

import (
	"net/http"
	"flag"
)

func main() {
	flag.Parse()

	logInfof("spr_api version:%s started...\n", VERSION)
	http.HandleFunc("/servicelist", servicelist)
	http.HandleFunc("/serverlist", serverlist)
//	http.HandleFunc("/createnode", createnode)
	http.HandleFunc("/addservice", addservice)
	http.HandleFunc("/delservice", delservice)
	http.HandleFunc("/addserver", addserver)
	http.HandleFunc("/delserver", delserver)

	err := http.ListenAndServe(LISTEN, nil)
	if err != nil {
		logInfof("ListenAndServe:", err)
	}

}

