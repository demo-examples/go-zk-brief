package main

import (
	"fmt"
	"net/http"
	"log"
//	"github.com/samuel/go-zookeeper/zk"
)

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
}

func main() {

	http.HandleFunc("/get", get)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}


