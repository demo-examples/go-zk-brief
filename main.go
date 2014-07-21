package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {

	fmt.Println("pri_api started...")
	http.HandleFunc("/servicelist", servicelist)
	http.HandleFunc("/serverlist", serverlist)
//	http.HandleFunc("/createnode", createnode)
	http.HandleFunc("/addservice", addservice)
	http.HandleFunc("/delservice", delservice)
	http.HandleFunc("/addserver", addserver)
	http.HandleFunc("/delserver", delserver)

	err := http.ListenAndServe(LISTEN, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

