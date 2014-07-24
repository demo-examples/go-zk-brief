package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
)


func addserver(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addserver start...")
	defer handleError(w)

	r.ParseMultipartForm(DEFAULT_MIN_MEMORY)

	keys := r.Form["key"]
	destNames := r.Form["destName"]
	zkidcs := r.Form["zkidc"]
	serverHost := r.Form["serverHost"]
	serverPort := r.Form["serverPort"]

	fmt.Println(keys)

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

