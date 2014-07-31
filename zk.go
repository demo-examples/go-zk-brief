package main

import(
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

var (
	ZkConns  =  make(map[string] *zk.Conn, 3)
	err        error

	FlagSequence int32 = zk.FlagSequence
	DefaultACL = zk.WorldACL(0x1f)

)

func initZK() (error) {
	ZkConns[Conf.ZkQAName], _, err = zk.Connect([]string{Conf.ZkQAAddr}, Conf.ZkQATimeout)
	if err != nil {
		return err
	}

	ZkConns[Conf.ZkYZName], _, err = zk.Connect([]string{Conf.ZkYZAddr}, Conf.ZkYZTimeout)
	if err != nil {
		return err
	}

	ZkConns[Conf.ZkG1Name], _, err = zk.Connect([]string{Conf.ZkG1Addr}, Conf.ZkG1Timeout)
	if err != nil {
		return err
	}

	return nil
}


func connect(addr string, timeout time.Duration) (*zk.Conn, error) {
	zkConn, _, err := zk.Connect([]string{addr}, timeout)
	if(err != nil) {
		return nil, err
	}
	defer zkConn.Close()
	return zkConn, nil
}
