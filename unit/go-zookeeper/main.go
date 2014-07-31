package main

import(

	"fmt"
)

func main() {

	c := ZkConns["a"]
	children, stat, ch, err := c.ChildrenW("/soa")

	fmt.Print(children, stat, ch, err)
}
