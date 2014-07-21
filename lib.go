package main

import(
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"
)

func handleError(w http.ResponseWriter){

	if e:= recover(); e != nil {
//		fmt.Printf("%+s\n", e)
		var rtnError RtnError
		var rtnJson []byte
		rtnError.Code = 0
		rtnError.Reason = fmt.Sprintf("%+s", e)
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



