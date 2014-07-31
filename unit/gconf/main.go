package main

import(
	"fmt"
	"time"
	"github.com/Terry-Mao/goconf"
)

type Config struct {
	HttpBind           string     `goconf:"core:http.bind"`
	Key                string     `goconf:"core:key"`
	ZkPath             string     `goconf:"core:path"`

	ZkQAAddr           string          `goconf:"zkQA:addr"`
	ZkQATimeout        time.Duration   `goconf:"zkQA:timeout:time"`
	ZkQAName           string          `goconf:"zkQA:name"`

	ZkYZAddr           string          `goconf:"zkYZ:addr"`
	ZkYZTimeout        time.Duration   `goconf:"zkYZ:timeout:time"`
	ZkYZName           string          `goconf:"zkYZ:name"`

	ZkG1Addr           string          `goconf:"zkG1:addr"`
	ZkG1Timeout        time.Duration   `goconf:"zkG1:timeout:time"`
	ZkG1Name           string          `goconf:"zkG1:name"`

}


func main() {
	confFile := "./config.conf"
	gconf := goconf.New()
	if err := gconf.Parse(confFile); err != nil {
		panic(err)
	}

	Conf := &Config {
		HttpBind:     ":9999",
		Key:          "1122-3333",
		ZkPath:       "/soa/services2",
		ZkQAName:     "abc",
		ZkQATimeout:  time.Second *2,
	}

    core := gconf.Get("core")                                                   
	fmt.Printf("core:%v\n", core)
	zkQa := gconf.Get("zkQA")
	fmt.Printf("zkQA:%v\n", zkQa)
	zkYZ := gconf.Get("zkYZ")
	fmt.Printf("zkYZ:%v\n", zkYZ)


	fmt.Printf("%v\n", Conf)
	if err := gconf.Unmarshal(Conf); err != nil {
		panic(err)
	}
	
	fmt.Printf("%v\n", Conf)
}
