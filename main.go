package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
)

type ZkServer struct {
	ServiceEndpoint ServerConf
	additionalEndpoints interface{}   // not used
	status interface{}   // not used
	shard interface{}   // not used
}

type RtnError struct {
	Code int
	Reason interface{}
}

type ServerConf struct {
	Host string
	Port int
}

type Service struct {
	Service string
}

type RtnServicelist struct {
	Code int
	Services []Service
}


type RtnServerlist struct {
	Code int
	Servers []ServerConf
}

const (
	KEY = "1122-3434"
//	ZKHOST = "192.168.35.141"
	ZKHOST = "192.168.129.213"
//	ZKHOST = "192.168.113.212"
	ZKPORT = 2181
	ZKPATH = "/soa/services"
	ZKTIMEOUT = time.Second
)

func serverlist(w http.ResponseWriter, r *http.Request) {

	defer handleError(w)
	r.ParseForm()
	
	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkNodes := r.Form["zkNode"]

	var rtnError RtnError
//	var rtnServers RtnServerlist
//	var rtnJson []byte
	rtnError.Code = 0

	// 参数检验
	if(len(destNames) != 1 || len(zkNodes) != 1 || len(keys) !=1) {
		panic("wrong param num")
	}

	// 判断key是否正确
	if(keys[0] != KEY) {  // @todo 修改为通过私钥判断的?
		panic("wrong key")
	}

//	zkNode = zkNodes[0]
	fmt.Println("connect zk!")

	c, _, err := zk.Connect([]string{fmt.Sprintf("%s:%d", ZKHOST, ZKPORT)}, time.Second)
	if(err != nil) {
		panic(err)
	}
	defer c.Close()

	zkServerPath := ZKPATH + "/" + destNames[0]
	children, _, _, err := c.ChildrenW(zkServerPath)
	if(err != nil) {
		panic(err)
	}

	var servers []ServerConf
	for _, child := range children {
		jsonServer, _, err := c.Get(zkServerPath + "/" + child)
		if(err != nil) {
			panic(err)
		}
		var zkserver ZkServer

		err = json.Unmarshal(jsonServer, &zkserver)
		if(err != nil) {
			panic(err)
		}

		servers = append(servers, zkserver.ServiceEndpoint)
	}

	rtnServer := &RtnServerlist {
		Code : 1,
		Servers : servers,
	}

	jsonRtn, err := json.Marshal(rtnServer)
	if(err != nil) {
		panic(err)
	}
	fmt.Fprintf(w, string(jsonRtn))
}

func servicelist(w http.ResponseWriter, r *http.Request) {
	defer handleError(w)
	r.ParseForm()

	keys := r.Form["key"]
	zkNodes := r.Form["zkNode"]

	var rtnError RtnError
	var rtnJson []byte
	rtnError.Code = 0

	// 参数检验
	if(len(zkNodes) != 1 || len(keys) !=1) {
		panic("wrong param num")
	}

	// 判断key是否正确
	if(keys[0] != KEY) {  // @todo 修改为通过私钥判断的?
		panic("wrong key")
	}

//	zkNode = zkNodes[0]
//	fmt.Println("connect zk!")
	c, _, err := zk.Connect([]string{fmt.Sprintf("%s:%d", ZKHOST, ZKPORT)}, time.Second)
	if(err != nil) {
		panic(err)
	}
	defer c.Close()

	children, stat, ch, err := c.ChildrenW(ZKPATH)
	if err != nil {
		panic(err)
	}

	var services []Service

	for _, v := range children {
//		fmt.Println(i, v)
		services = append(services, Service{Service : v})
	}

	fmt.Println(services)

	rtnServices := &RtnServicelist{
		Code : 1,
		Services : services,
	}
	rtnJson, _ = json.Marshal(rtnServices)


	fmt.Printf("-----%+v %+v   %+v\n", children, stat, ch)
//	e := <-ch


	fmt.Fprintf(w, string(rtnJson))

}

func createnode(w http.ResponseWriter, r *http.Request) {
	defer handleError(w)
	r.ParseForm()

	keys := r.Form["keys"]
	destNames := r.Form["destName"]
	zkNodes := r.Form["zkNode"]

	// 参数检验
	if !checkParams(keys, destNames, zkNodes) {
		panic("wrong param num")
	}

	// 判断key是否正确
	if !checkKeys(keys[0]) {  // @todo 修改为通过私钥判断的?
		panic("wrong key")
	}

	c, _, err := zk.Connect([]string{fmt.Sprintf("%s:%d", ZKHOST, ZKPORT)}, ZKTIMEOUT)
	if err != nil {
		panic(err)
	}
	defer c.Close()


}


func main() {

	http.HandleFunc("/servicelist", servicelist)
	http.HandleFunc("/serverlist", serverlist)
	http.HandleFunc("/createnode", createnode)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}


func handleError(w http.ResponseWriter){

	if e:= recover(); e != nil {
//		fmt.Println(e)
		var rtnError RtnError
		var rtnJson []byte
		rtnError.Code = 0
		rtnError.Reason = e
		rtnJson, _ = json.Marshal(rtnError)
		fmt.Fprintf(w, string(rtnJson))
	}

}

// 检查key是否正确，正确返回true, 错误返回false
func checkKeys(key string) bool{
	if key == KEY {
		return true
	}
	return false
}

func checkParams(params ...[]string) bool{
	for _, param := range params {
		if len(param) != 1 {
			return false
		}
	}
	return true
}
