package main

import (
	"fmt"
	"strings"
)

func StringChallenge(str string) string {

	fmt.Println("the index ", strings.LastIndex(str, "div>a") )
	arrStr := strings.Split(str, "<")
	for _, v := range arrStr {
		fmt.Println(v)
		arrStr1 := strings.SplitAfter(v, ">")
		for _, v1 := range arrStr1 {
			fmt.Println("inner")
			fmt.Println(v1)
			fmt.Println("end of inner")

		}
	}
	return ""
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(StringChallenge("<div>abc</div><p><em><i>test test test</b></em></p>"))

}