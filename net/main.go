package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	netIfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	fmt.Println(PrettyPrint(netIfaces))
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
