package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"strconv"
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
)

type ZkServer struct {
	ServiceEndpoint ServerConf     `json:"serviceEndpoint"`
	AdditionalEndpoints struct{}   `json:"additionalEndpoints"`// not used
	Status interface{}             `json:status`  // not used
	Shard interface{}              `json:shard`   // not used
}

type RtnNormal struct {
	Code int      `json:"code"`
}

type RtnError struct {
	Code int               `json:"code"`
	Reason interface{}     `json:"reason"`
}

type ServerConf struct {
	Host string            `json: "host"`
	Port int               `json: "port"`
}

type ServerConf2 struct {
	Host string     `json:"host"`
	Port int        `json:"port"`
	Key string      `json:"key"`
}

type Service struct {
	Service string      `json:"service"`
}

type RtnServicelist struct {
	Code int                `json:"code"`
	Services []Service      `json:"services"`
}


type RtnServerlist struct {
	Code int                   `json:"code"`
	Servers []ServerConf2      `json:"servers"`
}



const (
	KEY = "1122-3434"
	ZKPATH = "/soa/services"
	ZKTIMEOUT = time.Second
)

//	ZKHOST = "192.168.35.141"
//	ZKHOST = "192.168.129.213"
//	ZKHOST = "192.168.113.212"
//	ZKPORT = 2181

var ZKHOST map[string] string= map[string] string {
	"qa" : "192.168.35.141:2181",
//	"qa" : "192.168.129.213:2182",
	"yz" : "192.168.129.213:2181",
	"g1" : "192.168.129.213:2181",
}

func serverlist(w http.ResponseWriter, r *http.Request) {

	defer handleError(w)
	r.ParseForm()
	
	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkidcs := r.Form["zkidc"]

	var rtnError RtnError
//	var rtnServers RtnServerlist
//	var rtnJson []byte
	rtnError.Code = 0

	// 参数检验
	checkParams(destNames, zkidcs, keys)
	// 判断key是否正确
	checkKeys(keys[0])

//	zkidc = zkidcs[0]
	fmt.Println("connect zk!")

	c, _, err := zk.Connect([]string{ZKHOST[zkidcs[0]]}, time.Second)
	if(err != nil) {
		panic(err)
	}
	defer c.Close()

	zkServerPath := ZKPATH + "/" + destNames[0]
	children, _, _, err := c.ChildrenW(zkServerPath)
	if(err != nil) {
		panic(err)
	}

	var servers []ServerConf2
	for _, child := range children {
		fmt.Println(child)
		jsonServer, _, err := c.Get(zkServerPath + "/" + child)
		if(err != nil) {
			panic(err)
		}
		var zkserver ZkServer

		err = json.Unmarshal(jsonServer, &zkserver)
		if(err != nil) {
			panic(err)
		}

		server := ServerConf2 {
			Host : zkserver.ServiceEndpoint.Host,
			Port : zkserver.ServiceEndpoint.Port,
			Key : child,
		}

		servers = append(servers, server)
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
	zkidcs := r.Form["zkidc"]

	var rtnError RtnError
	var rtnJson []byte
	rtnError.Code = 0

	// 参数检验
	checkParams(zkidcs, keys)
	// 判断key是否正确
	checkKeys(keys[0])

//	zkidc = zkidcs[0]
//	fmt.Println("connect zk!")
	c, _, err := zk.Connect([]string{ZKHOST[zkidcs[0]]}, time.Second)
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

func addservice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createnode start...")
	defer handleError(w)
	r.ParseForm()

	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkidcs := r.Form["zkidc"]

	// 参数检验
	checkParams(keys, destNames, zkidcs)
	// 判断key是否正确
	checkKeys(keys[0])

	c, _, err := zk.Connect([]string{ZKHOST[zkidcs[0]]}, ZKTIMEOUT)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	zkServerPath := ZKPATH + "/" + destNames[0]
//	zkServerPath := "/soa/services"

//	fmt.Println(zkServerPath)
//	resPath, err := c.Create(zkServerPath, []byte{}, 0, zk.WorldACL(0x1f))
//	fmt.Println(resPath)

	_, err = c.Create(zkServerPath, []byte{}, 0, zk.WorldACL(0x1f))
	if err != nil {
		panic(err)
	}
	rtnNormal := &RtnNormal {
		Code : 1,
	}
	rtnJson, _ := json.Marshal(rtnNormal)
	fmt.Fprint(w, string(rtnJson))
}

func delservice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delservice start...")
	defer handleError(w)
	r.ParseForm()

	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkidcs := r.Form["zkidc"]

	// 参数检验
	checkParams(keys, destNames, zkidcs)
	// 判断key是否正确
	checkKeys(keys[0])

	c, _, err := zk.Connect([]string{ZKHOST[zkidcs[0]]}, ZKTIMEOUT)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	zkServerPath := ZKPATH + "/" + destNames[0]

//	fmt.Println(zkServerPath)
	err = c.Delete(zkServerPath, -1)

	if err != nil {
		panic(err)
	}
	rtnNormal := &RtnNormal {
		Code : 1,
	}
	rtnJson, _ := json.Marshal(rtnNormal)
	fmt.Fprint(w, string(rtnJson))

}



func addserver(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addserver start...")
	defer handleError(w)
	r.ParseForm()

	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkidcs := r.Form["zkidc"]
	serverHost := r.Form["serverHost"]
	serverPort := r.Form["serverPort"]


	// 参数检验
	checkParams(keys, destNames, zkidcs, serverHost, serverPort)
	// 判断key是否正确
	checkKeys(keys[0])

	c, _, err := zk.Connect([]string{ZKHOST[zkidcs[0]]}, ZKTIMEOUT)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	zkServerPath := ZKPATH + "/" + destNames[0] + "/member_"
	fmt.Println(zkServerPath)

	serverValue := getServerValue(serverHost[0], serverPort[0])
//	serverValue := []byte{96, 97}
//	fmt.Println(serverValue)
//	_, err = c.Create(zkServerPath, serverValue, zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(0x1f))
//	zkPath, err := c.Create(zkServerPath, serverValue, zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(0x1f))
	zkPath, err := c.Create(zkServerPath, serverValue, zk.FlagSequence, zk.WorldACL(0x1f))
	fmt.Println(zkPath)
	if err != nil {
		panic(err)
	}
	rtnNormal := &RtnNormal {
		Code : 1,
	}
	rtnJson, _ := json.Marshal(rtnNormal)
	fmt.Fprint(w, string(rtnJson))
}


func delserver(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delserver start...")
	defer handleError(w)
	r.ParseForm()

	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkidcs := r.Form["zkidc"]
	serverKey := r.Form["serverKey"]

	// 参数检验
	checkParams(keys, destNames, zkidcs, serverKey)
	// 判断key是否正确
	checkKeys(keys[0])

	c, _, err := zk.Connect([]string{ZKHOST[zkidcs[0]]}, ZKTIMEOUT)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	zkServerPath := ZKPATH + "/" + destNames[0] + "/" + serverKey[0]

//	fmt.Println(zkServerPath)
	err = c.Delete(zkServerPath, -1)

	if err != nil {
		panic(err)
	}
	rtnNormal := &RtnNormal {
		Code : 1,
	}
	rtnJson, _ := json.Marshal(rtnNormal)
	fmt.Fprint(w, string(rtnJson))

}




func main() {

	http.HandleFunc("/servicelist", servicelist)
	http.HandleFunc("/serverlist", serverlist)
//	http.HandleFunc("/createnode", createnode)
	http.HandleFunc("/addservice", addservice)
	http.HandleFunc("/delservice", delservice)
	http.HandleFunc("/addserver", addserver)
	http.HandleFunc("/delserver", delserver)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}


func handleError(w http.ResponseWriter){

	if e:= recover(); e != nil {
//		fmt.Printf("%+s\n", e)
		var rtnError RtnError
		var rtnJson []byte
		rtnError.Code = 0
		rtnError.Reason = 		fmt.Sprintf("%+s", e)
		rtnJson, _ = json.Marshal(rtnError)
//		fmt.Println(string(rtnJson))
		fmt.Fprintf(w, string(rtnJson))
	}

}

// 检查key是否正确，正确返回true, 错误返回false
func checkKeys(key string) {
	if key == KEY {
		return
	}
	panic("wrong keys")
}

func checkParams(params ...[]string) {
	for _, param := range params {
		if len(param) != 1 {
			panic("wrong params")
		}
	}
	return
}

func getServerValue(host string, strPort string) []byte{
	port, err := strconv.Atoi(strPort)
	if err!= nil {
		panic("port should be int")
	}
	
	serverConf := ServerConf {
		Host : host,
		Port : port,
	}
	var nothing struct{}
	serverValue := &ZkServer {
		ServiceEndpoint : serverConf,
		AdditionalEndpoints : nothing,
		Status : "ALIVE",
		Shard : 1,
	}
//	fmt.Println(serverValue)
	valueJson, err := json.Marshal(serverValue)
//	fmt.Println(string(valueJson), err)
	if err != nil {
		panic(err)
	}
	return valueJson
}



