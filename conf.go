package main

import (
	"time"
)

const (
	LISTEN = ":9090"
	KEY = "1122-3434"
	ZKPATH = "/soa/services"
	ZKTIMEOUT = time.Second
)

//	ZKHOST = "192.168.35.141"
//	ZKHOST = "192.168.129.213"
//	ZKHOST = "192.168.113.212"
//	ZKPORT = 2181

var ZKHOST map[string] string= map[string] string {
	"qa" : "192.168.35.141:2181",
//	"qa" : "192.168.129.213:2182",
	"yz" : "192.168.129.213:2181",
	"g1" : "192.168.129.213:2181",
}


