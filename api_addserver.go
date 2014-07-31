package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)


func addserver(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(DEFAULT_MIN_MEMORY)

	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkidcs := r.Form["zkidc"]
	serverHost := r.Form["serverHost"]
	serverPort := r.Form["serverPort"]

	input := fmt.Sprintf("keys:%v, zkidcs:%v, destNames:%v, serverHost:%v, serverPort:%v", keys, zkidcs, destNames, serverHost, serverPort)
	api := "addserver"
	defer handleError(w, input, api)



	// 参数检验
	checkParams(keys, destNames, zkidcs, serverHost, serverPort)
	// 判断key是否正确
	checkKeys(keys[0])

	c := ZkConns[zkidcs[0]]

	zkServerPath := Conf.ZkPath + "/" + destNames[0] + "/smember_"

	serverValue := getServerValue(serverHost[0], serverPort[0])

	_, err = c.Create(zkServerPath, serverValue, FlagSequence, DefaultACL)

	if err != nil {
		panic(err)
	}
	rtnNormal := &RtnNormal {
		Code : 1,
	}
	rtnJson, _ := json.Marshal(rtnNormal)

	rtnStr := string(rtnJson)
	fmt.Fprintf(w, rtnStr)

	apilog(input, api, rtnStr)   // 日志记录

}

