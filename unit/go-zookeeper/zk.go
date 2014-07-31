package main

import(
	"github.com/samuel/go-zookeeper/zk"
	"time"

)

var (
	ZkConns  =  make(map[string] *zk.Conn, 3)
	err error
)

func init() {
	ZkConns["a"], _, err = zk.Connect([]string{"192.168.35.141:2181"}, time.Second)
}
