package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
)


func serverlist(w http.ResponseWriter, r *http.Request) {

	defer handleError(w)
	r.ParseForm()
	
	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkidcs := r.Form["zkidc"]

	fmt.Println(keys, destNames, zkidcs)

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

	c, _, err := zk.Connect([]string{ZKHOST[zkidcs[0]]}, ZKTIMEOUT)
	if(err != nil) {
		panic(err)
	}
	defer c.Close()

	zkServerPath := ZKPATH + "/" + destNames[0]
	children, _, err := c.Children(zkServerPath)
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

