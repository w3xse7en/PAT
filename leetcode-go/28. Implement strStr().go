package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("hello", ""))
}
