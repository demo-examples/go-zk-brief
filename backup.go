package main

import (
	"io/ioutil"
	"time"
	"fmt"
	"github.com/golang/glog"
)

func backup() {
	glog.Infof("backup start interval: %s", Conf.BkInterval)
	ticker := time.NewTicker(Conf.BkInterval)

	for t := range ticker.C { // 定时循环
		now := t.Format("20060102_150405")
		glog.Infof("backup zk in: %v, t:%v", now, t)

		for k, c := range ZkConns { // 每个idc
			path := fmt.Sprintf("%s/zkbkg_%s_%s.zk.bkg", Conf.BkPath, k, now)
			var content string
			children, _, err := c.Children(Conf.ZkPath)
			if err != nil {
				glog.Errorf("error in bakup zk, err:%s", err)
			}

			for _, child := range children {   // 每个服务
				content = fmt.Sprintf("%s[%s]\n", content, child)
				zkServerPath := Conf.ZkPath + "/" + child
				children, _, err := c.Children(zkServerPath)
				if err != nil {
					glog.Errorf("error in bakup zk, err2:%s", err)
				}
				for _, child := range children {  // 每个服务下面的结点
					jsonServer, _, err := c.Get(zkServerPath + "/" + child)

					if err != nil {
						glog.Errorf("error in bakup zk, err2:%s", err)
					}

					content = fmt.Sprintf("%s%s\n", content, jsonServer)
				}
			}
			ioutil.WriteFile(path, []byte(content), 0755)
		}
	}
}
