package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
)


func servicelist(w http.ResponseWriter, r *http.Request) {
	defer handleError(w)
	r.ParseForm()

	keys := r.Form["key"]
	zkidcs := r.Form["zkidc"]

	fmt.Println(keys, zkidcs)

	var rtnError RtnError
	var rtnJson []byte
	rtnError.Code = 0

	// 参数检验
	checkParams(zkidcs, keys)
	// 判断key是否正确
	checkKeys(keys[0])

//	zkidc = zkidcs[0]
//	fmt.Println("connect zk!")
	c, _, err := zk.Connect([]string{ZKHOST[zkidcs[0]]}, ZKTIMEOUT)
	if(err != nil) {
		panic(err)
	}
	defer c.Close()

	children, stat, err := c.Children(ZKPATH)
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


	fmt.Printf("-----%+v %+v   \n", children, stat)
//	e := <-ch


	fmt.Fprintf(w, string(rtnJson))

}


