package main

import (
	"time"
	"flag"
	"github.com/Terry-Mao/goconf"
)

var(
	Conf         *Config
	confFile     string
)

type Config struct {
	HttpBind           string     `goconf:"core:http.bind"`
	Key                string     `goconf:"core:key"`
	ZkPath             string     `goconf:"core:zkpath"`

	ZkQAAddr           string          `goconf:"zkQA:addr"`
	ZkQATimeout        time.Duration   `goconf:"zkQA:timeout:time"`
	ZkQAName           string          `goconf:"zkQA:name"`

	ZkYZAddr           string          `goconf:"zkYZ:addr"`
	ZkYZTimeout        time.Duration   `goconf:"zkYZ:timeout:time"`
	ZkYZName           string          `goconf:"zkYZ:name"`

	ZkG1Addr           string          `goconf:"zkG1:addr"`
	ZkG1Timeout        time.Duration   `goconf:"zkG1:timeout:time"`
	ZkG1Name           string          `goconf:"zkG1:name"`

	BkPath             string          `goconf:"backup:path"`
	BkInterval         time.Duration   `goconf:"backup:interval:time"`

}

const (
	VERSION = "0.2.0"

	DEFAULT_MIN_MEMORY = 32 << 20
	DEFAULT_MAX_MEMORY = 1024

	ZKPREFIX = "smember_"
)

func init() {
	flag.StringVar(&confFile, "c", "config.conf", "set config file path")
}

func InitConfig() error {
	gconf := goconf.New()
	if err := gconf.Parse(confFile); err != nil {
		return err
	}

	Conf = &Config {
		HttpBind:     ":9999",
		Key:          "1122-3333",
		ZkPath:       "/soa/services2",
		ZkQAName:     "abc",
	}
//	debugf("%v\n", Conf)
	if err := gconf.Unmarshal(Conf); err != nil {
		return err
	}
//	debugf("%v\n", Conf)
	return nil

}




