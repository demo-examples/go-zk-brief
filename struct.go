package main



type ZkServer struct {
	ServiceEndpoint ServerConf     `json:"serviceEndpoint"`
	AdditionalEndpoints struct{}   `json:"additionalEndpoints"`// not used
	Status interface{}             `json:status`  // not used
	Shard interface{}              `json:shard`   // not used
}

type RtnNormal struct {
	Code int      `json:"code"`
}

type RtnError struct {
	Code int               `json:"code"`
	Reason interface{}     `json:"reason"`
}

type ServerConf struct {
	Host string            `json: "host"`
	Port int               `json: "port"`
}

type ServerConf2 struct {
	Host     string     `json:"host"`
	Port     int        `json:"port"`
	Key      string     `json:"key"`
	Readonly bool       `json:"readonly"`
}

type Service struct {
	Service string      `json:"service"`
}

type RtnServicelist struct {
	Code int                `json:"code"`
	Services []Service      `json:"services"`
}


type RtnServerlist struct {
	Code int                   `json:"code"`
	Servers []ServerConf2      `json:"servers"`
}

