package main

import(
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/golang/glog"
)

type ApiError struct{
	Api   string
	Input string
	Err   interface{}
}

func init() {
}

func (e *ApiError) Error() string {
	rtn := fmt.Sprintf("[%s]\tparams:%s\t%v", e.Api, e.Input, e.Err)
	return rtn
}

func handleError(w http.ResponseWriter, input string, api string){

	if e:= recover(); e != nil {
		var rtnError RtnError
		var rtnJson []byte
		rtnError.Code = 0
		rtnError.Reason = fmt.Sprintf("%+s", e)
		rtnJson, _ = json.Marshal(rtnError)

		fmt.Fprintf(w, string(rtnJson))   // 返回值

		apiErr := &ApiError{api, input, e}
		glog.Error(apiErr.Error())
	}

}

/*
func handleError(w http.ResponseWriter){

	if e:= recover(); e != nil {
//		fmt.Printf("%+s\n", e)
		var rtnError RtnError
		var rtnJson []byte
		rtnError.Code = 0
		rtnError.Reason = fmt.Sprintf("%+s", e)
		rtnJson, _ = json.Marshal(rtnError)
		fmt.Fprintf(w, string(rtnJson))   // 返回值

		error(e.Error())   // 打印错误日志

	}

}
*/

