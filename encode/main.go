package main

import (
	"encoding/base64"
	"fmt"
)

var xorb = []byte{0xf3, 0x45, 0xdd, 0x21, 0x72, 0x86, 0x93}

func Encode(data string) string {
	dataBytes := []byte(data)
	for i := range dataBytes {
		dataBytes[i] = dataBytes[i] ^ xorb[i%len(xorb)]
	}
	return base64.StdEncoding.EncodeToString(dataBytes)

}

func decode(data string) string {
	dataBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}

	for i := range dataBytes {
		dataBytes[i] = dataBytes[i] ^ xorb[i%len(xorb)]
	}
	return string(dataBytes)

}


func main() {
	fmt.Println(Encode("testingfileretrieval"))

	fmt.Println(Encode("uNuiVj0qTB6qnSTHPbogO95nvG2CicpyIShCRa0hLNOi3vLzPzxLYdArJgb6zWo/avCC8GxizyLN4fe0+BIG7Q=="))
}
