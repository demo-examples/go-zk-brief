package main

import(
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
)


func serverlist(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkidcs := r.Form["zkidc"]

	input := fmt.Sprintf("keys:%v, zkidcs:%v, destName:%v", keys, zkidcs, destNames)
	api := "serverlist"
	defer handleError(w, input, api)

	var rtnError RtnError
	rtnError.Code = 0

	// 参数检验
	checkParams(destNames, zkidcs, keys)
	// 判断key是否正确
	checkKeys(keys[0])

	c := ZkConns[zkidcs[0]]

	zkServerPath := Conf.ZkPath + "/" + destNames[0]
	children, _, err := c.Children(zkServerPath)
	if(err != nil) {
		panic(err)
	}

	var servers []ServerConf2
	for _, child := range children {
		debug(child)
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
			Host     : zkserver.ServiceEndpoint.Host,
			Port     : zkserver.ServiceEndpoint.Port,
			Key      : child,
			Readonly : !strings.HasPrefix(child, ZKPREFIX),
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
	rtnStr := string(jsonRtn)
	fmt.Fprintf(w, rtnStr)

	apilog(input, api, rtnStr)   // 日志记录
}

