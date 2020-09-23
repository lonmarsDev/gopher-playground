package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("https://gmail.googleapis.com/gmail/v1/users/105818564067529236868/messages?access_token=ya29.a0AfH6SMD2aYuiofdBxPrvqET-OS0qLIS1As_WabwgvpUnaaSr4utU_VF_YQHIHaHs1wcuo6bU5fR1TX2dG0uD-aRe5EPYuFxn40ZT5kSm8smoEMRz4L2CESt91H_qyUTFOsiTVIXrUqP4G5ck7rgkQQ2e3OPWdHkgNP0")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

}
