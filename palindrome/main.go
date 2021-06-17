package main

import "fmt"

func main() {
	word := "aniani"
	fmt.Println(isPalindrome(word))
}

func reverse(s string) string {
	r := []rune(s)
	//for _, v := range r {
	//
	//}

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindrome(word string) bool {
	if word == reverse(word) {
		return true
	}
	return false
}