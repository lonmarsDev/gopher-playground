package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"

)

func main() {


	client := &http.Client{ }

	req, err := http.NewRequest("GET", "https://api.elasticemail.com//v2/contact/list?apikey=908E50PO9IKLKLKPOKIJUHY76JUH76FHLMNN9087UJ82BAF4CFB1CCE26AA738EFF510FA91EE7C18A974C9982C0DD109D0", nil)
	// ...
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var byteData  []byte
	_, err = resp.Body.Read(byteData)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	b1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(b1))


	fmt.Printf( resp.Status )
	fmt.Println( resp.ContentLength )
	// ...


	resp2, err := http.Get("https://api.elasticemail.com//v2/contact/list?apikey=908E50PO9IKLKLKPOKIJUHY76JUH76FHLMNN9087UJ82BAF4CFB1CCE26AA738EFF510FA91EE7C18A974C9982C0DD109D0")
	if err != nil {
		// handle error
	}
	defer resp2.Body.Close()

	b, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(b))
	// ...

}
