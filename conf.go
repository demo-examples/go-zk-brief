package main

import (
	"time"
	"github.com/Terry-Mao/goconf"
)

var(
	Conf *config
)

type Config struct {
	HttpBind           []string     `goconf:"core:http.bind:,"`
	Key                []string     `goconf:"core:key"`
	ZkPath             []string     `goconf:"zk:path:,"`
	ZkTimeout          []string     `goconf:"zk.timeout:,"`
	ZkAddr             []string     `goconf:"zk.addr:,"`
	
}

const (
	VERSION = "0.1.0"

	DEFAULT_MIN_MEMORY = 32 << 20
	DEFAULT_MAX_MEMORY = 1024

)

//	ZKHOST = "192.168.35.141"
//	ZKHOST = "192.168.129.213"
//	ZKHOST = "192.168.113.212"
//	ZKPORT = 2181

var ZKHOST map[string] string= map[string] string {
	"test" : "192.168.35.141:2181",
//	"qa" : "192.168.35.141:2181",
	"qa" : "192.168.129.213:2181",
	"yz" : "yz-log-master-02:2181",
	"g1" : "g1-cdc-wrk-02:2181",
}


func init() {
	flag.StringVar(&confFile, "c", "config.conf", "set config file path")
}

func InitConfig() error {
	gconf := goconf.New()
	if err := goconf.Parse(confFile); err != nil {
		panic(fmt.Sprint("when parsing the file:%v, export:%v", confFile, err))
	}

	Conf = &Config {
		HttpBind:     ":9090",
		Key:          "1122-3434",
		ZkPath:       ""
	}
}
