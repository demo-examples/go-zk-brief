package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
)

type RtnError struct {
	Code int
	Reason string
}

type ServerConf struct {
	Host string
	Port int
}

type RtnGet struct {
	Code int
	Data []ServerConf
}

const (
	KEY = "1122-3434"
//	ZKHOST = "192.168.35.141"
	ZKHOST = "192.168.129.213"
	ZKPORT = 2181
	ZKPATH = "/soa/services"
)

var c interface{}

func get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkHosts := r.Form["zkHost"]

	var rtnError RtnError
	var rtnGet RtnGet
	var rtnJson []byte
	rtnError.Code = 0

	// 参数检验
	if(len(destNames) !=1 || len(zkHosts) != 1 || len(keys) !=1) {
		rtnError.Reason = "wrong param num"
		rtnJson, _ = json.Marshal(rtnError)
		fmt.Fprintf(w, string(rtnJson))
		return
	}

	// 判断key是否正确
	if(keys[0] != KEY) {  // @todo 修改为通过私钥判断的?
		rtnError.Reason="wrong key"
		rtnJson, _ = json.Marshal(rtnError)
		fmt.Fprintf(w, string(rtnJson))
		return
	}

//	destName = destNames[0]
//	zkHost = zkHosts[0]

	connectZK()

	fmt.Println(rtnGet)
	fmt.Fprintf(w, "111")
}

func connectZK() {
	fmt.Println("connect zk!")
	c, _, err := zk.Connect([]string{fmt.Sprintf("%s:%d", ZKHOST, ZKPORT)}, time.Second)
	if(err != nil) {
		panic(err)
	}
	defer c.Close()

	children, stat, ch, err := c.ChildrenW(ZKPATH)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-----%+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("=====%+v\n", e)
}

func main() {

	http.HandleFunc("/get", get)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}


