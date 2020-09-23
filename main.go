package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const token string = "" // Define token

func main() {
	client := http.Client{}
	resp, err := client.Get("https://gmail.googleapis.com/gmail/v1/users/105818564067529236868/messages?access_token=" + token)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

}
