package main

import (
	"time"
)

const (
	VERSION = "0.1.0"

	LISTEN = ":9090"            // 应用启动端口
	KEY = "1122-3434"           // 验证码
	ZKPATH = "/soa/services"    // 基于此目录下
	ZKTIMEOUT = time.Second     // zk连接超时时间

	DEFAULT_MIN_MEMORY = 32 << 20
	DEFAULT_MAX_MEMORY = 1024

)

//	ZKHOST = "192.168.35.141"
//	ZKHOST = "192.168.129.213"
//	ZKHOST = "192.168.113.212"
//	ZKPORT = 2181

var ZKHOST map[string] string= map[string] string {
	"test" : "192.168.35.141:2181",
	"qa" : "192.168.35.141:2181",
//	"qa" : "192.168.129.213:2182",
	"yz" : "192.168.129.213:2181",
	"g1" : "192.168.129.213:2181",
}


