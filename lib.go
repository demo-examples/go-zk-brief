package main

import(
	"strconv"
	"strings"
	"encoding/json"
)

// 检查服务类型是否以"smember_"开头
func checkServerKey(serverKey string) {
	if strings.HasPrefix(serverKey, ZKPREFIX) {
		return
	}
	panic("auth refused: server type")
}

// 检查key是否正确，正确返回true, 错误返回false
func checkKeys(key string) {
	if key == Conf.Key {
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
//	debug(serverValue)
	valueJson, err := json.Marshal(serverValue)
//	debug(string(valueJson))
	if err != nil {
		panic(err)
	}
	return valueJson
}



