package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
)


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

