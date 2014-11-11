package main

import(
	"fmt"
	"sort"
	"net/http"
	"encoding/json"
)


func servicelist(w http.ResponseWriter, r *http.Request) {
	debug("servicelist")
	r.ParseForm()

	keys := r.Form["key"]
	zkidcs := r.Form["zkidc"]

	input := fmt.Sprintf("keys:%v, zkidcs:%v", keys, zkidcs)
	api := "servicelist"
	defer handleError(w, input, api)

	var rtnError RtnError
	var rtnJson []byte
	rtnError.Code = 0

	// 参数检验
	checkParams(zkidcs, keys)
	// 判断key是否正确
	checkKeys(keys[0])
	debug(zkidcs[0])

	c := ZkConns[zkidcs[0]]
	debugf("[%T]%v\n", c, c)

	children, _, err := c.Children(Conf.ZkPath)

	debugf("[%T]%v\n", children, children)

	if err != nil {
		panic(err)
	}
	var services []Service

	for _, v := range children {
//		fmt.Println(i, v)
		services = append(services, Service{Service : v})
	}

//	fmt.Println(services)
	sort.Stable(ByServiceName(services))
	rtnServices := &RtnServicelist{
		Code : 1,
		Services : services,
	}
	rtnJson, _ = json.Marshal(rtnServices)
	rtnStr := string(rtnJson)

	fmt.Fprintf(w, rtnStr)

	apilog(input, api, rtnStr)   // 日志记录

}

type ByServiceName []Service


func (p ByServiceName) Len() int    { return len(p) }
func (p ByServiceName) Swap(i, j int)  { p[i], p[j] = p[j], p[i] }
func (p ByServiceName) Less(i, j int) bool { return p[i].Service < p[j].Service }

