package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	datas, _ := getDisabledNetInterfaceWin()
	for _, v := range datas {
		fmt.Println("debbb")
		fmt.Println(v)
		fmt.Println("debbb-ed")

	}

	// fmt.Println(data)
	// fmt.Println(len(data))
}

func getDisabledNetInterfaceWin() ([]string, error) {

	cmd := exec.Command("c:\\Windows\\System32\\wbem\\wmic.exe", "nic", "where", "NetConnectionStatus=0", "get", "NetConnectionID")
	result, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	out := strings.Trim(string(result), " \n\t\r")
	arrs := strings.Split(out, "\n")
	var interfaces []string
	for _, v := range arrs {
		if strings.TrimSpace(v) != "NetConnectionID" || strings.TrimSpace(v) != "" {
			interfaces = append(interfaces, v)
			// fmt.Println("-----damay------------")
			// fmt.Println(v)
			// fmt.Println("---------damy-end--------")
		}
	}

	return interfaces, nil
}
