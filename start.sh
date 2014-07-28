#!/bin/sh

rm ./go-zk-brief
go build github.com/zhaoweiguo/go-zk-brief
./go-zk-brief -v=1 -log_dir="/tmp/go-zk-brief/" -stderrthreshold=FATAL
#nohup ./go-zk-brief -v=1 -log_dir="/tmp/go-zk-brief/" -stderrthreshold=FATAL &



